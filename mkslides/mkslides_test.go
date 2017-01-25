//
// mkslides.go - A simple command line utility that uses Markdown
// to generate a sequence of HTML5 pages that can be used for presentations.
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
package mkslides

import (
	"strings"
	"testing"
	"text/template"
)

func TestBasic(t *testing.T) {
	src := `

# Opening Slide

## This is an introduction

+ me, <me@example.org>
+ [mygitrepo.org](http://mygitrepo.org)
+ February 31, 2016, Somebody's University, Someplace

--

# Slide Two

| col A | col B | Col C |
|-------|-------|-------|
| A     | One   | 1     |

--

## Slide Three

This is slide three, just a random paragraph of text. Blah, blah, blah, blah, blah, blah, boink!


`

	tmpl, err := template.New("test.tmpl").Parse(DefaultTemplateSource)
	if err != nil {
		t.Errorf("Can't parse DefaultTemplateSource templates %s", err)
		t.FailNow()
	}

	titles := []string{
		"<h1>Opening Slide</h1>",
		"<h1>Slide Two</h1>",
		"<h2>Slide Three</h2>",
	}

	slides := MarkdownToSlides("test.html", "This is just a test", "", "", []byte(src))
	if len(slides) != 3 {
		t.Errorf("Was expected three slides %+v\n", slides)
	}

	for i, slide := range slides {
		s, err := MakeSlideString(tmpl, slide)
		if err != nil {
			t.Errorf("MakeSlideString() failed %d - %s", i, err)
		}
		if strings.Contains(s, titles[i]) == false {
			t.Errorf("Expected %q in slide %d -> %s", titles[i], i, s)
		}
	}
}
