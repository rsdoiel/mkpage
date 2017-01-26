//
// mkslides.go - A simple command line utility that uses Markdown to generate a sequence of HTML5 pages that can be used for presentations.
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
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// My packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/mkpage"
)

var (
	usage = `USAGE: %s [OPTIONS] [FILENAME]`

	description = `
SYNOPSIS

%s creates slides from a single Markdown file delimiting
each side by a "---". 
"---" 
`

	examples = ``

	// Standard Options
	showHelp    bool
	showVersion bool
	showLicense bool

	// Application options
	presentationTitle string
	cssPath           string
	jsPath            string
	templateFName     string
	templateSource    = mkpage.DefaultSlideTemplateSource
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.StringVar(&presentationTitle, "t", "", "Presentation title")
	flag.StringVar(&presentationTitle, "title", "", "Presentation title")
	flag.StringVar(&cssPath, "c", cssPath, "Specify the CSS file to use")
	flag.StringVar(&cssPath, "css", cssPath, "Specify the CSS file to use")
	flag.StringVar(&jsPath, "j", jsPath, "Specify a js file to include")
	flag.StringVar(&jsPath, "js", jsPath, "Specify a js file to include")
	flag.StringVar(&templateFName, "template", templateFName, "Specify an HTML template to use")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	cfg := cli.New(appName, "MKPAGE", fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version), mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	//cfg.ExampleText = examples

	if showHelp == true {
		fmt.Println(cfg.Usage())
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	if templateFName != "" {
		src, err := ioutil.ReadFile(templateFName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n", templateFName, err)
			os.Exit(1)
		}
		templateSource = string(src)
	}

	//FIXME: If it is markdown file then assign fname that value, otherwise it's a template add it to the
	// list of templates to compile.
	var fname string
	args := flag.Args()
	if len(args) > 0 {
		//fname, args = args[0], args[1:]
		fname = args[0]
	}
	if fname == "" {
		fmt.Fprintf(os.Stderr, "Missing filename")
		os.Exit(1)
	}
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fname = strings.TrimSuffix(path.Base(fname), path.Ext(fname))
	tmpl, err := template.New("slide").Parse(templateSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Build the slides
	slides := mkpage.MarkdownToSlides(fname, presentationTitle, cssPath, jsPath, src)
	// Render the slides
	for i, slide := range slides {
		err := mkpage.MakeSlideFile(tmpl, slide)
		if err == nil {
			// Note: Give some feed back when slide written successful
			fmt.Fprintf(os.Stdout, "Wrote %02d-%s.html\n", slide.CurNo, slide.FName)
		} else {
			// Note: Display an error if we have a problem
			fmt.Fprintf(os.Stderr, "Can't process slide %d, %s\n", i, err)
		}
	}
	// Render the TOC slide
	slide, err := mkpage.SlidesToTOCSlide(slides)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create a table of contents slide", err)
		os.Exit(1)
	}
	err = mkpage.MakeTOCSlideFile(tmpl, slide)
	if err == nil {
		// Note: Give some feed back when slide written successful
		fmt.Fprintf(os.Stdout, "Wrote toc-%s.html\n", slide.FName)
	} else {
		// Note: Display an error if we have a problem
		fmt.Fprintf(os.Stderr, "Can't write toc-%s.html, %s\n", slide.FName, err)
	}
}
