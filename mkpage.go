//
// mkpage is an experiment in a light weight template and markdown processor.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright 2017 R. S. Doiel
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
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	// 3rd Party Packages
	"github.com/russross/blackfriday"
)

const (
	// Version of the mkpage package.
	Version = "v0.0.13"

	// LicenseText for cli programs
	LicenseText = `
%s %s

Copyright 2017 R. S. Doiel

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

	// SOMEDAY: should add XML, BibTeX, YaML support...

	// The default HTML provided by md2slides package, you probably want to override this...
	DefaultTemplateSource = `<!DOCTYPE html>
<html>
<head>
  {{if .title -}}<title>{{- .title -}}</title>{{- end}}
  {{if .csspath -}}<link href="{{ .csspath }}" rel="stylesheet" />{{else -}}
  {{with .mdfile -}}<link href="{{- . -}}" rel="application/markdown" />{{- end}}
  <style>
    body {
      width: 100%;
      height: 100%;
      margin: 0;
      padding: 0;
      font-size: calc(1em+1vw);
      font-family: sans-serif;
    }

    ul {
      list-style: disc;
      text-indent: 0.25em;
    }

    nav {
      position: absolute;
      top: 0em; 
      margin:0;
      padding:0.24em;
      width: 100%;
      height: 4em;
      text-align: left;
      font-size: 60%;
    }

	header {
		width: 100%;
		height: auto;
	}

	footer {
		width: 100%;
		height: auto;
	}

    section {
      width: 80%;
      height: auto;
	  padding: 5%;
    }
  </style>
  {{- end }}
</head>
<body>
  {{if .header -}}
  <header>{{- .header -}}</header>
  {{end}}
  {{if .nav -}}
  <nav>{{- .nav -}}</nav>
  {{end}}
  {{if .content -}}
  <section>{{ .content }}</section>
  {{end}}
  {{if .footer -}}
  <footer>{{ .footer }}</footer>
  {{end}}
</body>
</html>`
)

// ResolveData takes a data map and reads in the files and URL sources
// as needed turning the data into strings to be applied to the template.
func ResolveData(data map[string]string) (map[string]interface{}, error) {
	var out map[string]interface{}

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
			out[key] = string(blackfriday.MarkdownCommon([]byte(strings.TrimPrefix(val, MarkdownPrefix))))
		case strings.HasPrefix(val, JSONPrefix) == true:
			var o interface{}
			err := json.Unmarshal(bytes.TrimPrefix([]byte(val), []byte(JSONPrefix)), &o)
			if err != nil {
				return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
			}
			out[key] = o

		case strings.HasPrefix(val, "http://") == true || strings.HasPrefix(val, "https://") == true:
			resp, err := http.Get(val)
			if err != nil {
				return out, fmt.Errorf("Error from %s, %s", val, err)
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
							return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
						}
						out[key] = o
					case isContentType(contentTypes, "text/markdown") == true:
						out[key] = string(blackfriday.MarkdownCommon(buf))
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
				return out, fmt.Errorf("Can't read %s, %s", val, err)
			}
			ext := path.Ext(val)
			switch {
			case strings.Compare(ext, ".md") == 0:
				out[key] = string(blackfriday.MarkdownCommon(buf))
			case strings.Compare(ext, ".json") == 0:
				var o interface{}
				err := json.Unmarshal(buf, &o)
				if err != nil {
					return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
				}
				out[key] = o
			default:
				out[key] = string(buf)
			}
		}
	}
	return out, nil
}

// MakePage applies the key/value map to the template and renders to writer and returns an error if something goes wrong
func MakePage(wr io.Writer, tmpl *template.Template, keyValues map[string]string) error {
	data, err := ResolveData(keyValues)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	return tmpl.Execute(wr, data)
}

// MakePageString applies the key/value map to the template and renders the results to a string and error if someting goes wrong
func MakePageString(tmpl *template.Template, keyValues map[string]string) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakePage(wr, tmpl, keyValues)
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
	return strings.Join(result, sep)
}

//
// mkslide code
//

// Slide is the metadata about a slide to be generated.
type Slide struct {
	XMLName xml.Name `json:"-"`
	CurNo   int      `xml:"CurNo" json:"CurNo"`
	PrevNo  int      `xml:"PrevNo" json:"PrevNo"`
	NextNo  int      `xml:"NextNo" json:"NextNo"`
	FirstNo int      `xml:"FirstNo" json:"FirstNo"`
	LastNo  int      `xml:"LastNo" json:"LastNo"`
	FName   string   `xml:"FName" json:"FName"`
	Title   string   `xml:"Title" json:"Title"`
	Heading string   `xml:"Heading" json:"Heading"`
	Content string   `xml:"Content" json:"Content"`
	CSSPath string   `xml:"CSSPath,omitempty" json:"CSSPath,omitempty"`
	JSPath  string   `xml:"JSPath,omitempty" json:"JSPath,omitempty"`
}

var (
	// The default HTML provided by mkslides package, you probably want to override this...
	DefaultSlideTemplateSource = `<!DOCTYPE html>
<html>
<head>
    {{if .Title -}}<title>{{- .Title -}}</title>{{- end}}
    {{if .CSSPath -}}
<link href="{{ .CSSPath }}" rel="stylesheet" />
   {{else -}}
<style>
    body {
        width: 100%;
        height: 100%;
        margin: 10%;
        padding: 0;
        font-size: calc(2em+1vw);
        font-family: sans-serif;
    }
    
    ul {
        list-style: disc;
        text-indent: 0.25em;
    }
    
    nav {
        position: absolute;
        top: 0em; 
        margin:0;
        padding:0.24em;
        width: 100%;
        height: 4em;
        text-align: left;
        font-size: 60%;
    }
    
    section {
        width: 100%;
        height: auto;
    }
</style>
{{- end }}
</head>
<body>
    <nav>
{{ if ne .CurNo .FirstNo -}}
<a id="start-slide" href="{{printf "%02d-%s.html" .FirstNo .FName}}">Home</a>
{{- end}}
{{ if gt .CurNo .FirstNo -}} 
<a id="prev-slide" href="{{printf "%02d-%s.html" .PrevNo .FName}}">Prev</a>
{{- end}}
{{ if lt .CurNo .LastNo -}} 
<a id="next-slide" href="{{printf "%02d-%s.html" .NextNo .FName}}">Next</a>
{{- end}}
    </nav>
    <section>{{ .Content }}</section>
{{with .JSPath}}<script src="{{.}}"></script>{{end}}
<script>
(function (document, window) {
    'use strict';
    var start = document.getElementById('start-slide'),
        prev = document.getElementById('prev-slide'),
        next = document.getElementById('next-slide');
    
    
    document.onkeydown = function(e) {
        switch (e.keyCode) {
            /* case 32: */
            case 37:
            // Previous: left arrow
                if (prev) {
                    prev.click();
                }
                break;
            case 39:
                // Next: right arrow
                if (next) {
                    next.click();
                }
                break;
            case 72:
            case 83:
                // Home/Start: h, s
                if (start) {
                    start.click();
                }
                break;
        }
    };
}(document, window));
</script>
</body>
</html>
`
)

// MarkdownToSlides turns a markdown file into one or more Slide using the fname, title and cssPath provided
func MarkdownToSlides(fname string, title string, cssPath string, jsPath string, src []byte) []*Slide {
	var slides []*Slide

	// Note: handle legacy CR/LF endings as well as normal LF line endings
	if bytes.Contains(src, []byte("\r\n")) {
		src = bytes.Replace(src, []byte("\r\n"), []byte("\n"), -1)
	}

	// Note: We're only spliting on line that contains only "--",
	mdSlides := bytes.Split(src, []byte("\n--\n"))

	lastSlide := len(mdSlides) - 1
	for i, s := range mdSlides {
		//Note: Collect first heading for TOC slide
		heading := []byte("")
		hIndex := bytes.Index(s, []byte("# "))
		if hIndex > -1 {
			hEOL := bytes.Index(s[hIndex:], []byte("\n"))
			if hEOL > -1 {
				heading = s[hIndex : hIndex+hEOL]
			}
		}
		// Note: Convert slide's Markdown to HTML
		data := blackfriday.MarkdownCommon(s)
		// Assemble Slide
		slides = append(slides, &Slide{
			FName:   fname,
			CurNo:   i,
			PrevNo:  (i - 1),
			NextNo:  (i + 1),
			FirstNo: 0,
			LastNo:  lastSlide,
			Title:   title,
			Heading: string(bytes.TrimPrefix(heading, []byte("# "))),
			Content: string(data),
			CSSPath: cssPath,
			JSPath:  jsPath,
		})
	}
	return slides
}

// MakeSlide this takes a io.Writer, a template and slide and executes the template.
func MakeSlide(wr io.Writer, tmpl *template.Template, slide *Slide) error {
	return tmpl.Execute(wr, slide)
}

// MakeSlideFile this takes a template and slide and renders the results to a file.
func MakeSlideFile(tmpl *template.Template, slide *Slide) error {
	sname := fmt.Sprintf(`%02d-%s.html`, slide.CurNo, slide.FName)
	fp, err := os.Create(sname)
	if err != nil {
		return fmt.Errorf("%s %s\n", sname, err)
	}
	defer fp.Close()
	err = MakeSlide(fp, tmpl, slide)
	if err != nil {
		return fmt.Errorf("%s %s", sname, err)
	}
	return nil
}

// MakeSlideString this takes a template and slide and renders the results to a string
func MakeSlideString(tmpl *template.Template, slide *Slide) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakeSlide(wr, tmpl, slide)
	return buf.String(), err
}

// SlidesToTOCSlide takes an array of slide generating a new Slide structure
// whos content is a table of contents of other slides and their first headings.
func SlidesToTOCSlide(slides []*Slide) (*Slide, error) {
	var buf bytes.Buffer
	src := `
<h1>Table of contents</h1>
<ul>
{{range $slide := . -}}
	<li><a href="{{- printf "%02d-%s.html" $slide.CurNo $slide.FName -}}">{{- printf "%d &mdash; %s" $slide.CurNo $slide.Heading -}}</a></li>
{{- end}}
</ul>
`
	tmpl, err := template.New("slide").Parse(src)
	if err != nil {
		return nil, err
	}
	wr := io.Writer(&buf)
	err = tmpl.Execute(wr, slides)
	if err != nil {
		return nil, err
	}

	tocSlide := new(Slide)
	tocSlide.FName = slides[0].FName
	tocSlide.Title = slides[0].Title
	tocSlide.Content = buf.String()

	return tocSlide, nil
}

// MakeTOCSlideFile this takes a template and slide and renders the results to a file.
func MakeTOCSlideFile(tmpl *template.Template, slide *Slide) error {
	sname := fmt.Sprintf(`toc-%s.html`, slide.FName)
	fp, err := os.Create(sname)
	if err != nil {
		return fmt.Errorf("%s\n", sname, err)
	}
	defer fp.Close()
	err = MakeSlide(fp, tmpl, slide)
	if err != nil {
		return fmt.Errorf("%s", sname, err)
	}
	return nil
}
