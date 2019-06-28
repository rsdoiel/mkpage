//
// Package mkpage is an experiment in a light weight template and markdown processor.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2019, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package mkpage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	// 3rd Party Packages
	"github.com/BurntSushi/toml"
	"github.com/ghodss/yaml"
	"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	/*
		"github.com/mmarkdown/mmark/mast"
		"github.com/mmarkdown/mmark/mparser"
		"github.com/mmarkdown/mmark/render/man"
		mmarkout "github.com/mmarkdown/mmark/render/markdown"
		"github.com/mmarkdown/mmark/render/mhtml"
		"github.com/mmarkdown/mmark/render/xml"
		"github.com/mmarkdown/mmark/render/xml2"
	*/
	"github.com/rsdoiel/fountain"
	// FIXME: Should this be depreciated? duplicative of gomarkdown
	"gopkg.in/russross/blackfriday.v2"
)

const (
	Version = `v0.0.29`

	// LicenseText provides a string template for rendering cli license info
	LicenseText = `
%s %s

Copyright (c) 2019, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

	// Prefix for explicit string types

	// JSONPrefix designates a string as JSON formatted content
	JSONPrefix = "json:"
	// MarkdownPrefix designates a string as Markdown content
	MarkdownPrefix = "markdown:"
	// TextPrefix designates a string as text/plain not needed processing
	TextPrefix = "text:"
	// FountainPrefex designates a string as Fountain formatted content
	FountainPrefix = "fountain:"

	// SOMEDAY: should add XML, BibTeX, YaML support...

	// DateExp is the default format used by mkpage utilities for date exp
	DateExp = `[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]`
	// BylineExp is the default format used by mkpage utilities
	BylineExp = (`^[B|b]y\s+(\w|\s|.)+` + DateExp + "$")
	// TitleExp is the default format used by mkpage utilities
	TitleExp = `^#\s+(\w|\s|.)+$`

	// Front Matter Types
	FrontMatterIsUnknown = iota
	FrontMatterIsYAML
	FrontMatterIsTOML
	FrontMatterIsJSON
)

var (
	// DefaultTemplateSource is defined in init by Defaults["/templates/page.tmpl"]
	// Defaults is a map to assets defined in assets.go which is build with pkgasset and
	// the contents of the defaults folder in this repository.
	DefaultTemplateSource string

	// DefaultSlideTemplateSource provides the default HTML template for mkslides package,
	// you probably want to override this... is defined in init by Defaults["/templates/slides.tmpl"]
	// Defaults is a map to assets defined in assets.go which is build with pkgasset and
	// the contents of the defaults folder in this repository.
	DefaultSlideTemplateSource string
)

// SplitFronMatter takes a []byte input splits it into front matter
// source and Markdown source. If either is missing an empty []byte
// is returned for the missing element.
func SplitFrontMatter(input []byte) (int, []byte, []byte) {
	// YAML front matter uses ---
	if bytes.HasPrefix(input, []byte("---\n")) {
		parts := bytes.SplitN(bytes.TrimPrefix(input, []byte("---\n")), []byte("\n---\n"), 2)
		return FrontMatterIsYAML, parts[0], parts[1]
	}
	if bytes.HasPrefix(input, []byte("+++\n")) {
		parts := bytes.SplitN(bytes.TrimPrefix(input, []byte("+++\n")), []byte("\n+++\n"), 2)
		return FrontMatterIsTOML, parts[0], parts[1]
	}
	if bytes.HasPrefix(input, []byte("{\n")) {
		parts := bytes.SplitN(bytes.TrimPrefix(input, []byte("{\n")), []byte("\n}\n"), 2)
		src := []byte(fmt.Sprintf("{\n%s\n}\n", parts[0]))
		return FrontMatterIsJSON, src, parts[1]
	}
	// Handle case of no front matter
	return FrontMatterIsUnknown, []byte(""), input
}

// processorConfig takes front matter and returns
// a map[string]interface{}{} containing configuration for a target
// markup engine.
func processorConfig(frontMatterType int, frontMatterSrc []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	// Convert Front Matter to JSON
	switch frontMatterType {
	case FrontMatterIsYAML:
		// YAML Front Matter
		jsonSrc, err := yaml.YAMLToJSON(frontMatterSrc)
		if err != nil {
			return nil, fmt.Errorf("Can't parse YAML front matter, %s", err)
		}
		if err = json.Unmarshal(jsonSrc, &m); err != nil {
			return nil, err
		}
	case FrontMatterIsTOML:
		// TOML Front Matter
		if _, err := toml.Decode(fmt.Sprintf("%s", frontMatterSrc), &m); err != nil {
			return nil, err
		}
	case FrontMatterIsJSON:
		// JSON Front Matter
		if err := json.Unmarshal(frontMatterSrc, &m); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown front matter format")
	}
	return m, nil
}

// blackfridayExtensions takes a config (map[string]interface{}) and
// returns the ORed extenions flags for map["blackfriday"].
func blackfridayExtensions(config map[string]interface{}) blackfriday.Extensions {
	ext := blackfriday.NoExtensions
	if thing, ok := config["blackfriday"]; ok == true {
		extensions := thing.(map[string]interface{})
		for k, v := range extensions {
			onoff := false
			option := k
			switch v.(type) {
			case int:
				onoff = v.(int) == 1
			case int64:
				onoff = v.(int64) == 1
			case bool:
				onoff = v.(bool)
			case string:
				onoff = strings.ToLower(v.(string)) == "true"
			}
			if onoff {
				switch option {
				case "NoIntraEmphasis":
					ext |= blackfriday.NoIntraEmphasis
				case "Tables":
					ext |= blackfriday.Tables
				case "FencedCode":
					ext |= blackfriday.FencedCode
				case "Autolink":
					ext |= blackfriday.Autolink
				case "Strikethrough":
					ext |= blackfriday.Strikethrough
				case "LaxHTMLBlocks":
					ext |= blackfriday.LaxHTMLBlocks
				case "SpaceHeadings":
					ext |= blackfriday.SpaceHeadings
				case "HardLineBreak":
					ext |= blackfriday.HardLineBreak
				case "TabSizeEight":
					ext |= blackfriday.TabSizeEight
				case "Footnotes":
					ext |= blackfriday.Footnotes
				case "NoEmptyLineBeforeBlock":
					ext |= blackfriday.NoEmptyLineBeforeBlock
				case "HeadingIDs":
					ext |= blackfriday.HeadingIDs
				case "Titleblock":
					ext |= blackfriday.Titleblock
				case "AutoHeadingIDs":
					ext |= blackfriday.AutoHeadingIDs
				case "BackslashLineBreak":
					ext |= blackfriday.BackslashLineBreak
				case "DefinitionLists":
					ext |= blackfriday.DefinitionLists
				//case "CommonHTMLFlags":
				//	ext |= blackfriday.CommonHTMLFlags
				case "CommonExtensions":
					ext |= blackfriday.CommonExtensions
				}
			}
		}
	}
	return ext
}

// gomarkdownExtensions takes a config (map[string]interface{}) and
// returns the ORed extenions flags for map["markdown"].
func gomarkdownExtensions(config map[string]interface{}) parser.Extensions {
	ext := parser.NoExtensions
	if thing, ok := config["markdown"]; ok == true {
		extensions := thing.(map[string]interface{})
		for k, v := range extensions {
			onoff := false
			option := k
			switch v.(type) {
			case int:
				onoff = v.(int) == 1
			case int64:
				onoff = v.(int64) == 1
			case bool:
				onoff = v.(bool)
			case string:
				onoff = strings.ToLower(v.(string)) == "true"
			}
			fmt.Fprintf(os.Stderr, "DEBUG option %s value %t\n", option, onoff)
			if onoff {
				switch option {
				case "NoIntraEmphasis":
					ext |= parser.NoIntraEmphasis
				case "Tables":
					ext |= parser.Tables
				case "FencedCode":
					ext |= parser.FencedCode
				case "Autolink":
					ext |= parser.Autolink
				case "Strikethrough":
					ext |= parser.Strikethrough
				case "LaxHTMLBlocks":
					ext |= parser.LaxHTMLBlocks
				case "SpaceHeadings":
					ext |= parser.SpaceHeadings
				case "HardLineBreak":
					ext |= parser.HardLineBreak
				case "TabSizeEight":
					ext |= parser.TabSizeEight
				case "Footnotes":
					ext |= parser.Footnotes
				case "NoEmptyLineBeforeBlock":
					ext |= parser.NoEmptyLineBeforeBlock
				case "HeadingIDs":
					ext |= parser.HeadingIDs
				case "Titleblock":
					ext |= parser.Titleblock
				case "AutoHeadingIDs":
					ext |= parser.AutoHeadingIDs
				case "BackslashLineBreak":
					ext |= parser.BackslashLineBreak
				case "DefinitionLists":
					ext |= parser.DefinitionLists
				case "MathJax":
					ext |= parser.MathJax
				case "OrderedListStart":
					ext |= parser.OrderedListStart
				case "Attributes":
					ext |= parser.Attributes
				case "SuperSubscript":
					ext |= parser.SuperSubscript
				case "CommonExtensions":
					ext |= parser.CommonExtensions
				}
			}
		}
	}
	fmt.Fprintf(os.Stderr, "DEBUG extensions value %d\n", ext)
	return ext
}

// gomarkdownRenderOptions config (map[string]interface{}) and
// returns the ORed extenions flags for map["markdown"].
func gomarkdownRenderOptions(config map[string]interface{}) html.Flags {
	flag := html.FlagsNone
	if thing, ok := config["markdown"]; ok == true {
		flags := thing.(map[string]interface{})
		for k, v := range flags {
			option := k
			onoff := false
			switch v.(type) {
			case int:
				onoff = v.(int) == 1
			case int64:
				onoff = v.(int64) == 1
			case bool:
				onoff = v.(bool)
			case string:
				onoff = strings.ToLower(v.(string)) == "true"
			}
			fmt.Fprintf(os.Stderr, "DEBUG render option %s value %t\n", option, onoff)
			if onoff {
				switch option {
				case "SkipHTML":
					flag |= html.SkipHTML
				case "SkipImages":
					flag |= html.SkipImages
				case "SkipLinks":
					flag |= html.SkipLinks
				case "Safelink":
					flag |= html.Safelink
				case "NofollowLinks":
					flag |= html.NofollowLinks
				case "NoreferrerLinks":
					flag |= html.NoreferrerLinks
				case "HrefTargetBlank":
					flag |= html.HrefTargetBlank
				case "CompletePage":
					flag |= html.CompletePage
				case "UseXHTML":
					flag |= html.UseXHTML
				case "FootnoteReturnLinks":
					flag |= html.FootnoteReturnLinks
				case "FootnoteNoHRTag":
					flag |= html.FootnoteNoHRTag
				case "Smartypants":
					flag |= html.Smartypants
				case "SmartypantsFractions":
					flag |= html.SmartypantsFractions
				case "SmartypantsDashes":
					flag |= html.SmartypantsDashes
				case "SmartypantsLatexDashes":
					flag |= html.SmartypantsLatexDashes
				case "SmartypantsAngledQuotes":
					flag |= html.SmartypantsAngledQuotes
				case "SmartypantsQuotesNBSP":
					flag |= html.SmartypantsQuotesNBSP
				case "TOC":
					flag |= html.TOC
				case "CommonFlags":
					flag |= html.CommonFlags
				}
			}
		}
	}
	fmt.Fprintf(os.Stderr, "DEBUG render options value %d\n", flag)
	return flag
}

// markdownProcessor wraps gomarkdown, mmark, fountain and blackfriday.v2
// handling the splitting off the front matter if present and configuration
// via front matter.
func markdownProcessor(input []byte) ([]byte, error) {
	frontMatterType, frontMatterSrc, mdSrc := SplitFrontMatter(input)
	//NOTE: should inspect front matter for Markdown parsing configuration
	config, err := processorConfig(frontMatterType, frontMatterSrc)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stderr, "DEBUG config -> %+v\n", config)
	if thing, ok := config["markup"]; ok == true {
		markup := thing.(string)
		fmt.Fprintf(os.Stderr, "DEBUG markup is %s\n", markup)
		switch markup {
		case "markdown":
			//FIXME: Should this be the default processor?
			// Should I drop blackfriday?
			ext := gomarkdownExtensions(config)
			p := parser.NewWithExtensions(ext)

			htmlFlags := gomarkdownRenderOptions(config)
			opts := html.RendererOptions{Flags: htmlFlags}
			r := html.NewRenderer(opts)

			return markdown.ToHTML(mdSrc, p, r)
		case "fountain":
			return fountainProcessor(input)
		case "blackfriday":
			ext := blackfridayExtensions(config)
			//FIXME: Should support blackfriday.HTMLFlags too
			if ext == blackfriday.NoExtensions {
				return blackfriday.Run(mdSrc, blackfriday.WithNoExtensions()), nil
			} else {
				return blackfriday.Run(mdSrc, blackfriday.WithExtensions(ext)), nil
			}
		default:
			return nil, fmt.Errorf("unknown markup engine")
		}
	}
	// Default to gomarkdown markdown processor
	// with CommonExtensions and CommonHTMLFlags
	p := parser.NewWithExtensions(parser.CommonExtensions)
	opts := html.RendererOptions{Flags: html.CommonFlags}
	r := html.NewRenderer(opts)
	return markdown.ToHTML(mdSrc, p, r), nil

	// NOTE: Might want to have default as
	//return blackfriday.Run(mdSrc, blackfriday.WithExtensions(blackfriday.ComonExtensions|blackfriday.CommonHTMLFlags))
	// NOTE: Old default was
	//return blackfriday.Run(mdSrc)
}

// fountainProcessor wraps fountain.Run() splitting off the front
// matter if present.
func fountainProcessor(input []byte) ([]byte, error) {
	var err error

	frontMatterType, frontMatterSrc, fountainSrc := SplitFrontMatter(input)
	config, err := processorConfig(frontMatterType, frontMatterSrc)
	if err != nil {
		return nil, err
	}
	// Default Settings (override via front matter directives)
	fountain.AsHTMLPage = false
	fountain.InlineCSS = false
	fountain.LinkCSS = false
	fountain.IncludeCSS = ""
	if thing, ok := config["fountain"]; ok == true {
		cfg := thing.(map[string]interface{})
		for k, v := range cfg {
			switch v.(type) {
			case bool:
				onoff := v.(bool)
				switch k {
				case "AsHTMLPage":
					fountain.AsHTMLPage = onoff
				case "InlineCSS":
					fountain.InlineCSS = onoff
				case "LinkCSS":
					fountain.LinkCSS = onoff

				}
			case string:
				if k == "IncludeCSS" {
					fountain.IncludeCSS = v.(string)
				}
			default:
				return nil, fmt.Errorf("Unknown fountain option %q", k)
			}
		}
	}
	src, err := fountain.Run(fountainSrc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: %s\n", err)
	}
	return src
}

// ResolveData takes a data map and reads in the files and URL sources
// as needed turning the data into strings to be applied to the template.
func ResolveData(data map[string]string) (map[string]interface{}, error) {
	var (
		out map[string]interface{}
	)

	isContentType := func(vals []string, target string) bool {
		for _, h := range vals {
			if strings.Contains(h, target) == true {
				return true
			}
		}
		return false
	}

	out = make(map[string]interface{})
	for key, val := range data {
		switch {
		case strings.HasPrefix(val, TextPrefix) == true:
			out[key] = strings.TrimPrefix(val, TextPrefix)
		case strings.HasPrefix(val, MarkdownPrefix) == true:
			src, err := markdownProcessor([]byte(strings.TrimPrefix(val, MarkdownPrefix)))
			if err != nil {
				return out, err
			}
			out[key] = fmt.Sprintf("%s", src)
		case strings.HasPrefix(val, FountainPrefix) == true:
			src, err := fountainProcessor([]byte(strings.TrimPrefix(val, FountainPrefix)))
			if err != nil {
				return out, err
			}
			out[key] = fmt.Sprintf("%s", src)
		case strings.HasPrefix(val, JSONPrefix) == true:
			var o interface{}
			err := json.Unmarshal(bytes.TrimPrefix([]byte(val), []byte(JSONPrefix)), &o)
			if err != nil {
				return out, fmt.Errorf("Can't JSON decode (%s) %s, %s", key, val, err)
			}
			out[key] = o
		case strings.HasPrefix(val, "http://") == true || strings.HasPrefix(val, "https://") == true:
			resp, err := http.Get(val)
			if err != nil {
				return out, fmt.Errorf("Error from (%s) %s, %s", key, val, err)
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				buf, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return out, err
				}
				if contentTypes, ok := resp.Header["Content-Type"]; ok == true {
					switch {
					case isContentType(contentTypes, "application/json") == true:
						var o interface{}
						err := json.Unmarshal(buf, &o)
						if err != nil {
							return out, fmt.Errorf("Can't JSON decode (%s) %s, %s", key, val, err)
						}
						out[key] = o
					case isContentType(contentTypes, "text/markdown") == true:
						out[key] = string(markdownProcessor(buf))
					case isContentType(contentTypes, "text/fountain") == true:
						out[key] = string(fountainProcessor(buf))
					default:
						out[key] = string(buf)
					}
				} else {
					out[key] = string(buf)
				}
			}
		default:
			buf, err := ioutil.ReadFile(val)
			if err != nil {
				return out, fmt.Errorf("Can't read (%s) %q, %s", key, val, err)
			}
			ext := path.Ext(val)
			switch {
			case strings.Compare(ext, ".fountain") == 0 ||
				strings.Compare(ext, ".spmd") == 0:
				out[key] = string(fountainProcessor(buf))
			case strings.Compare(ext, ".md") == 0:
				out[key] = string(markdownProcessor(buf))
			case strings.Compare(ext, ".json") == 0:
				var o interface{}
				err := json.Unmarshal(buf, &o)
				if err != nil {
					return out, fmt.Errorf("Can't JSON decode (%s) %s, %s", key, val, err)
				}
				out[key] = o
			default:
				out[key] = string(buf)
			}
		}
	}
	return out, nil
}

// MakePage applies the key/value map to the named template in tmpl and renders to writer and returns an error if something goes wrong
func MakePage(wr io.Writer, templateName string, tmpl *template.Template, keyValues map[string]string) error {
	data, err := ResolveData(keyValues)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	return tmpl.ExecuteTemplate(wr, templateName, data)
}

// MakePageString applies the key/value map to the named template tmpl and renders the results to a string and error if someting goes wrong
func MakePageString(templateName string, tmpl *template.Template, keyValues map[string]string) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakePage(wr, templateName, tmpl, keyValues)
	return buf.String(), err
}

// RelativeDocPath calculate the relative path from source to target based on
// implied common base.
//
// Example:
//
//     docPath := "docs/chapter-01/lesson-02.html"
//     cssPath := "css/site.css"
//     fmt.Printf("<link href=%q>\n", MakeRelativePath(docPath, cssPath))
//
// Output:
//
//     <link href="../../css/site.css">
//
func RelativeDocPath(source, target string) string {
	var result []string

	sep := string(os.PathSeparator)
	dname, _ := path.Split(source)
	for i := 0; i < strings.Count(dname, sep); i++ {
		result = append(result, "..")
	}
	result = append(result, target)
	p := strings.Join(result, sep)
	if strings.HasSuffix(p, "/.") {
		return strings.TrimSuffix(p, ".")
	}
	return p
}

//
// Below is addition code to support mkslides
//

// Slide is the metadata about a slide to be generated.
type Slide struct {
	CurNo   int    `json:"cur_no,omitemtpy"`
	PrevNo  int    `json:"prev_no,omitempty"`
	NextNo  int    `json:"next_no,omitempty"`
	FirstNo int    `json:"first_no,omitempty"`
	LastNo  int    `json:"last_no,omitempty"`
	FName   string `json:"filename,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	CSSPath string `json:"csspath,omitempty"`
	JSPath  string `json:"jspath,omitempty"`
	CSS     string `json:"css,omitempty"`
	Header  string `json:"header,omitempty"`
	Footer  string `json:"footer,omitempty"`
	Nav     string `json:"nav,omitempty"`
}

// MarkdownToSlides turns a markdown file into one or more Slide structs
// Which populate predefined key/value pairs for later rendering in Markdown
func MarkdownToSlides(fname string, mdSource []byte) []*Slide {
	var slides []*Slide

	// Note: handle legacy CR/LF endings as well as normal LF line endings
	if bytes.Contains(mdSource, []byte("\r\n")) {
		mdSource = bytes.Replace(mdSource, []byte("\r\n"), []byte("\n"), -1)
	}
	// Note: We're only spliting on line that contain "--", not lines ending with where text might end with "--"
	mdSlides := bytes.Split(mdSource, []byte("\n--\n"))

	lastSlide := len(mdSlides) - 1
	for i, s := range mdSlides {
		slides = append(slides, &Slide{
			FName:   strings.TrimSuffix(path.Base(fname), path.Ext(fname)),
			CurNo:   i,
			PrevNo:  (i - 1),
			NextNo:  (i + 1),
			FirstNo: 0,
			LastNo:  lastSlide,
			Content: string(markdownProcessor(s)),
		})
	}
	return slides
}

// MakeSlide this takes a io.Writer, a template, key/value map pairs and Slide struct.
// It resolves the data int key/value pairs, merges the prefined mapping from Slide struct
// then executes the template.
func MakeSlide(wr io.Writer, templateName string, tmpl *template.Template, keyValues map[string]string, slide *Slide) error {
	data, err := ResolveData(keyValues)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	// Merge the slide metadata into data pairs for template
	data["filename"] = slide.FName
	data["cur_no"] = slide.CurNo
	data["prev_no"] = slide.PrevNo
	data["next_no"] = slide.NextNo
	data["first_no"] = slide.FirstNo
	data["last_no"] = slide.LastNo
	data["content"] = slide.Content
	data["header"] = slide.Header
	data["footer"] = slide.Header
	data["nav"] = slide.Nav
	return tmpl.ExecuteTemplate(wr, templateName, data)
}

// MakeSlideFile this takes a template and slide and renders the results to a file.
func MakeSlideFile(templateName string, tmpl *template.Template, keyValues map[string]string, slide *Slide) error {
	sname := fmt.Sprintf(`%02d-%s.html`, slide.CurNo, strings.TrimSuffix(path.Base(slide.FName), path.Ext(slide.FName)))
	fp, err := os.Create(sname)
	if err != nil {
		return fmt.Errorf("%s %s", sname, err)
	}
	defer fp.Close()
	err = MakeSlide(fp, templateName, tmpl, keyValues, slide)
	if err != nil {
		return fmt.Errorf("%s %s", sname, err)
	}
	return nil
}

// MakeSlideString this takes a template and slide and renders the results to a string
func MakeSlideString(templateName string, tmpl *template.Template, keyValues map[string]string, slide *Slide) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakeSlide(wr, templateName, tmpl, keyValues, slide)
	return buf.String(), err
}

// NormalizeDate takes a MySQL like date string and returns a time.Time or error
func NormalizeDate(s string) (time.Time, error) {
	switch len(s) {
	case len(`2006-01-02 15:04:05 -0700`):
		dt, err := time.Parse(`2006-01-02 15:04:05 -0700`, s)
		return dt, err
	case len(`2006-01-02 15:04:05`):
		dt, err := time.Parse(`2006-01-02 15:04:05`, s)
		return dt, err
	case len(`2006-01-02`):
		dt, err := time.Parse(`2006-01-02`, s)
		return dt, err
	default:
		return time.Time{}, fmt.Errorf("Can't format %s, expected format like 2006-01-02 15:04:05 -0700", s)
	}
}

// Walk takes a start path and walks the file system to process Markdown files for useful elements.
func Walk(startPath string, filterFn func(p string, info os.FileInfo) bool, outputFn func(s string, info os.FileInfo) error) error {
	err := filepath.Walk(startPath, func(p string, info os.FileInfo, err error) error {
		// Are we interested in this path?
		if filterFn(p, info) == true {
			// Yes, so send to output function.
			if err := outputFn(p, info); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// Grep looks for the first line matching the expression
// in src.
func Grep(exp string, src string) string {
	re, err := regexp.Compile(exp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q is not a valid, %s\n", exp, err)
		return ""
	}
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		s := re.FindString(line)
		if len(s) > 0 {
			return s
		}
	}
	return ""
}

func init() {
	if bString, ok := Defaults["/templates/page.tmpl"]; ok == true {
		DefaultTemplateSource = fmt.Sprintf("%s", bString)
	}
	if bString, ok := Defaults["/templates/slides.tmpl"]; ok == true {
		DefaultSlideTemplateSource = fmt.Sprintf("%s", bString)
	}
}
