//
// mkpage is an experiment in a light weight template and markdown processor.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
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
	"path"
	"strings"
	"text/template"

	// 3rd Party Packages
	"github.com/russross/blackfriday"
)

const (
	// Version of the mkpage package.
	Version = "v0.0.7"

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
  <style>
/**
 * site.css - stylesheet for the Library's Digital Library Development Group's sandbox.
 *
 * orange: #FF6E1E;
 *
 * Secondary pallet:
 *
 * lightgrey: #C8C8C8
 * grey: #76777B
 * darkgrey: #616265
 * slategrey: #AAA99F
 *
 * Impact Pallete see: http://identity.caltech.edu/web/colors
 */
body {
     margin: 0;
     border: 0;
     padding: 0;
     width: 100%;
     height: 100%;
     color: black;
     background-color: white;
     /*
     color: #FF6E1E;
     background-color: #AAA99F; /* #76777B;*/
     */
     font-family: Open Sans, Helvetica, Sans-Serif;
     font-size: 16px;
}

header {
     position: relative;
     display: block;
     color: white;
     background-color: white;
     z-index: 2;
     height: 105px;
     vertical-align: middle;
}

header img {
     position: relative;
     display: inline;
     padding-left: 20px;
     margin: 0;
     height: 42px;
     padding-top: 32px;
}

header h1 {
     position: relative;
     display: inline-block;
     margin: 0;
     border: 0;
     padding: 0;
     font-size: 3em;
     font-weight: normal;
     vertical-align: 0.78em;
     color: #FF6E1E;
}

header a, header a:link, header a:visited, header a:active, header a:hover, header a:focus {
     color: #FF6E1E;
     background-color: inherit;
}


a, a:link, a:visited {
     color: #76777B;
     background-color: inherit;
     text-decoration: none;
}

a:active, a:hover, a:focus {
    color: #FF6E1E;
    font-weight: bolder;
}

nav {
     position: relative;
     display: block;
     width: 100%;
     margin: 0;
     padding: 0;
     font-size: 0.78em;
     vertical-align: middle;
     color: black;
     background-color: #AAA99F; /* #76777B;*/
     text-align: left;
}

nav div {
     display: inline-block;
     /* padding-left: 10em; */
     margin-left: 10em;
     margin-right: 0;
}

nav a, nav a:link, nav a:visited, nav a:active {
     color: white;
     background-color: inherit;
     text-decoration: none;
}

nav a:hover, nav a:focus {
     color: #FF6E1E;
     background-color: inherit;
     text-decoration: none;
}


nav div h2 {
     position: relative;
     display: block;
     min-width: 20%;
     margin: 0;
     font-size: 1.24em;
     font-style: normal;
}

nav div > ul {
     display: none;
     padding-left: 0.24em;
     text-align: left;
}

nav ul {
     display: inline-block;
     padding-left: 0.24em;
     list-style-type: none;
     text-align: left;
     text-decoration: none; 
}

nav ul li {
     display: inline;
     padding: 1em;
}

section {
     position: relative;
     display: inline-block;
     width: 100%;
     min-height: 84%;
     color: black;
     background-color: white;
     margin: 0;
     padding-top 0;
     padding-bottom: 2em;
     padding-left: 1em;
     padding-right: 0;
}

section h1 {
    font-size: 1.32em;
}

section h2 {
    font-size: 1.12em;
    font-weight: italic;
}

section h3 {
    font-size: 1em;
    text-transform: uppercase;
}

section ul {
    display: block;
    list-style: inside;
    list-style-type: square;
    margin: 0;
    padding-left: 1.24em;
}

aside {
     margin: 0;
     border: 0;
     padding-left: 1em;
     position: relative;
     display: inline-block;
     text-align: right;
}

aside h2 {
     font-size: 1em;
     text-transform: uppercase;
}

aside h2 > a {
     font-style: normal;
}

aside ul {
     margin: 0;
     padding: 0;
     display: block;
     list-style-type: none;
}

aside ul li {
     font-size: 0.82em;
}

aside ul > ul {
     padding-left: 1em;
     font-size: 0.72em;
}

footer {
     position: fixed;
     bottom: 0;
     display: block;
     width: 100%;
     height: 2em;
     color: white;
     background-color: #616265;

     font-size: 80%;
     text-align: left;
     vertical-align: middle;
     z-index: 2;
}

footer h1, footer span, footer address {
     position: relative;
     display: inline-block;
     margin: 0;
     padding-left: 0.24em;
     font-family: Open Sans, Helvetica, Sans-Serif;
     font-size: 1em;
}

footer h1 {
     font-weight: normal;
}

footer a, footer a:link, footer a:visited, footer a:active, footer a:focus, footer a:hover {
     padding: 0;
     display: inline;
     margin: 0;
     color: #FF6E1E;
     text-decoration: none;
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
</html>
`
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
