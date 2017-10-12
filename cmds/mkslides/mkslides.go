//
// mkslides.go - A simple command line utility that uses Markdown
// to generate a sequence of HTML5 pages that can be used for presentations.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
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

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkpage"
	"github.com/caltechlibrary/tmplfn"
)

const (
	usage = `USAGE: %s [OPTIONS] [KEY/VALUE DATA PAIRS] MARKDOWN_FILE [TEMPLATE_FILENAMES]`

	description = `

SYNOPSIS

%s converts a Markdown file into a sequence of HTML5 slides using the
key/value pairs to populate the templates and render to stdout.

Features

+ Use Markdown to write your presentation in one file
+ Separate slides by "--" and a new line (e.g. \n versus \r\n)
+ Apply the default template or use your own
+ Control Layout and display with HTML5 and CSS

%s is based on mkpage with the difference that multiple pages
result from a single Markdown file. To manage the linkage between
slides some predefined template variables is used.

+ Title which would hold the page title for presentation
+ CSSPath which would hold the path to your CSS File.
+ Content holds the extracted for each slide
+ CurNo which holds the current page number
+ FirstNo which holds the first slide's page number (e.g. 00)
+ LastNo which holds the last slides page number (e..g length of slide deck minus one)
+ PrevNo which holds the previous slide number if CurNo is create than 0
+ NextNo which holds the next slide number if CurNo is not the last slide
+ FName is the filename for presentation

In your custom templates these should be exist to link everything together
as expected.  In addition you may want to include JavaScript to allow mapping
actions like "next slide" to the space bar or mourse click.

CONFIGURATION

+ MKPAGE_TEMPLATES - specify where to find the template(s) to use for slides

`

	examples = `

EXAMPLE

In this example we're using the default slide template.
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

If you saved this as presentation.md you can run the following
command to generate slides

    %s "Title=text:My Presentation" \
	    "CSSPath=text:css/slides.css" presentation.md

Using a custom template would look like

    %s -t custom-slides.tmpl \
        "Title=text:My Presentation" \
	    "CSSPath=text:css/slides.css" presentation.md

This would result in the following webpages

+ 00-presentation.html
+ 01-presentation.html
+ 02-presentation.html

`
)

var (
	// Standard Options
	showHelp     bool
	showVersion  bool
	showLicense  bool
	showExamples bool

	// Application Options
	cssPath           string
	jsPath            string
	mdFName           string
	presentationTitle string
	showTemplate      bool
	templateFNames    string
)

func init() {
	// Standard options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

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
	cfg := cli.New(appName, "MKPAGE", mkpage.Version)
	cfg.LicenseText = fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName, appName, appName, appName)

	// Process flags and update the environment as needed.
	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}

	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
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
		fmt.Println(mkpage.DefaultSlideTemplateSource)
		os.Exit(0)
	}

	// Make sure we have a configured command to run
	templateSources := []string{}
	templateFNames = cfg.MergeEnv("templates", templateFNames)
	if len(templateFNames) > 0 {
		for _, fname := range strings.Split(templateFNames, ":") {
			templateSources = append(templateSources, fname)
		}
	}

	data := map[string]string{}
	for i, arg := range args {
		switch {
		case strings.Contains(arg, "=") == true:
			// Update data map
			pair := strings.SplitN(arg, "=", 2)
			if len(pair) != 2 {
				fmt.Fprintf(os.Stderr, "Can't read pair (%d) %s\n", i+1, arg)
				os.Exit(1)
			}
			data[pair[0]] = pair[1]
		case path.Ext(arg) == ".md":
			mdFName = arg
		default:
			// Must be the template source
			templateSources = append(templateSources, arg)
		}
	}

	// Read in the Markdown file
	mdSrc, err := ioutil.ReadFile(mdFName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s, %s\n", mdFName, err)
		os.Exit(1)
	}

	// Default Template Name is slides.tmpl
	templateName := "slides.tmpl"

	// Create our Tmpl with its function map
	tmpl := tmplfn.New(tmplfn.AllFuncs())

	// Load ant user supplied templates
	if len(templateSources) > 0 {
		if err := tmpl.ReadFiles(templateSources...); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		templateName = templateSources[0]
	} else {
		// Read any templates from stdin that might be present
		if cli.IsPipe(os.Stdin) == true {
			buf, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			tmpl.Add(templateName, buf)
		} else {
			// Load our default template maps
			if err := tmpl.Add(templateName, mkpage.Defaults["/templates/slides.tmpl"]); err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
		}
	}

	// Assemble our templates
	t, err := tmpl.Assemble()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Build the slides
	slides := mkpage.MarkdownToSlides(mdFName, mdSrc)

	// Render the slides
	for i, slide := range slides {
		// Merge slide data with rest of command line map (e.g. "Title=text:My Presentation" "CSSPath=text:css/slides.css")
		err := mkpage.MakeSlideFile(templateName, t, data, slide)
		if err == nil {
			// Note: Give some feed back when slide written successful
			fmt.Fprintf(os.Stdout, "Wrote %02d-%s.html\n", slide.CurNo, strings.TrimSuffix(path.Base(slide.FName), path.Ext(slide.FName)))
		} else {
			// Note: Display an error if we have a problem
			fmt.Fprintf(os.Stderr, "Can't process slide %d, %s\n", i, err)
		}
	}
}
