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
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkslides"
	"github.com/caltechlibrary/tmplfn"
)

const (
	usage = `USAGE: %s [OPTIONS] [FILES]`

	description = `
SYNOPSIS

%s converts a Markdown file into a sequence of HTML5 slides.

+ Use Markdown to write your presentation in one file
+ Separate slides by "--" and a new line (e.g. \n versus \r\n)
+ Apply the simple default template or use your own
+ Control Layout and display with HTML5 and CSS

CONFIGURATION

+ MKSLIDES_CSS - specify the CSS file to include
+ MKSLIDES_JS - specify the JS file to include
+ MKSLIDES_MARKDOWN - the markdown file to process
+ MKSLIDES_PRESENTATION_TITLE - specify the title of the presentation
+ MKSLIDES_TEMPLATES - specify where to find the templates to use 
`
	examples = `
EXAMPLE

Here's an example of a three slide presentation

    Welcome to [%s](../)
    by R. S. Doiel, <rsdoiel@caltech.edu>

    --

    # %s

    %s can generate multiple HTML5 pages from
    one markdown file.  It splits the markdown file
    on each "--" 

    --

    Thank you

    Hope you enjoy [%s](https://github.com/caltechlbrary/%s)


If you save this as presentation.md and run "%s presentation.md" it would
generate the following webpages

+ 00-presentation.html
+ 01-presentation.html
+ 02-presentation.html

`
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	cssPath           string
	jsPath            string
	mdFName           string
	presentationTitle string
	showTemplate      bool
	templateFNames    string
	templateSource    = mkslides.DefaultTemplateSource
)

func init() {
	// Standard options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")

	// Application specific options
	flag.StringVar(&cssPath, "c", "", "Specify the CSS file to use")
	flag.StringVar(&cssPath, "css", "", "Specify the CSS file to use")
	flag.StringVar(&jsPath, "j", "", "Specify the JavaScript file to use")
	flag.StringVar(&jsPath, "js", "", "Specify the JavaScript file to use")
	flag.StringVar(&mdFName, "m", "", "Markdown filename")
	flag.StringVar(&mdFName, "markdown", "", "Markdown filename")
	flag.StringVar(&presentationTitle, "p", "", "Presentation title")
	flag.StringVar(&presentationTitle, "presentation-title", "", "Presentation title")
	flag.BoolVar(&showTemplate, "s", false, "display the default template")
	flag.BoolVar(&showTemplate, "show-template", false, "display the default template")
	flag.StringVar(&templateFNames, "t", "", "A colon delimited list of HTML templates to use")
	flag.StringVar(&templateFNames, "templates", "", "A colon delimited list of HTML templates to use")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configure app
	cfg := cli.New(appName, "MKSLIDES", fmt.Sprintf(mkslides.LicenseText, appName, mkslides.Version), mkslides.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName)

	// Process flags and update the environment as needed.
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

	if showTemplate == true {
		fmt.Println(mkslides.DefaultTemplateSource)
		os.Exit(0)
	}

	// Find the markdown/template filename if one is given on the command line
	for _, arg := range args {
		switch path.Ext(arg) {
		case ".md":
			mdFName = arg
		case ".css":
			cssPath = arg
		case ".js":
			jsPath = arg
		default:
			if len(templateFNames) > 0 {
				templateFNames = strings.Join([]string{templateFNames, arg}, ":")
			} else {
				templateFNames = arg
			}
		}
	}

	// Make sure we have a configured command to run after populating from command line args
	mdFName = cfg.CheckOption("markdown", cfg.MergeEnv("markdown", mdFName), true)
	templateFNames = cfg.MergeEnv("templates", templateFNames)
	cssPath = cfg.MergeEnv("css", cssPath)
	jsPath = cfg.MergeEnv("js", jsPath)

	// Read in the Markdown file
	mdSrc, err := ioutil.ReadFile(mdFName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s, %s\n", mdFName, err)
		os.Exit(1)
	}

	var (
		tmpl      *template.Template
		tmplFuncs = tmplfn.Join(tmplfn.TimeMap, tmplfn.PageMap)
	)

	//NOTE: If template is provided, read it in and replace templateSource content
	if len(templateFNames) > 0 {
		templateSources := strings.Split(templateFNames, ":")
		tmpl, err = tmplfn.Assemble(tmplFuncs, templateSources...)
	} else {
		tmpl, err = tmplfn.Assemble(tmplFuncs, templateSource)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Build the slides
	slides := mkslides.MarkdownToSlides(mdFName, presentationTitle, cssPath, jsPath, mdSrc)
	// Render the slides
	for i, slide := range slides {
		err := mkslides.MakeSlideFile(tmpl, slide)
		if err == nil {
			// Note: Give some feed back when slide written successful
			fmt.Fprintf(os.Stdout, "Wrote %02d-%s.html\n", slide.CurNo, strings.TrimSuffix(path.Base(slide.FName), path.Ext(slide.FName)))
		} else {
			// Note: Display an error if we have a problem
			fmt.Fprintf(os.Stderr, "Can't process slide %d, %s\n", i, err)
		}
	}
}
