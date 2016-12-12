//
// mkpage is a thought experiment in a light weight template and markdown processor
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
	"os"
	"path"
	"strings"
	"text/template"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkpage"
	"github.com/caltechlibrary/tmplfn"
)

var (
	usage = `USAGE: %s [OPTION] [KEY/VALUE DATA PAIRS] TEMPLATE_FILENAME [TEMPLATE_FILENAMES]`

	description = `
SYNOPSIS

Using the key/value pairs populate the template(s) and render to stdout.

CONFIGURATION

You can set a local default template path by using environment variables.

+ MKPAGE_TEMPLATES - is the colon delimited list of template paths
`

	examples = `
EXAMPLE

Template

    Date: {{- .now}}

    Hello {{.name -}},
    
    The current weather is

    {{.weather}}

    Thank you

	{{.signature}}

Render the template above (i.e. myformletter.template) would be accomplished from the following
data sources--

 + "now" and "name" are strings
 + "weatherForcast" comes from a URL
 + "license" comes from a file in our local disc

That would be expressed on the command line as follows

	%s "now=text:$(date)" "name=text:Little Frieda" \
		"weather=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" \
		signature=testdata/signature.txt \
		testdata/myformletter.template

Golang's text/template docs can be found at 

      https://golang.org/pkg/text/template/
`

	// Standard Options
	showHelp    bool
	showVersion bool
	showLicense bool

	// Application Options
	templateFNames string
	showTemplate   bool
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showLicense, "l", false, "show license")
	flag.BoolVar(&showLicense, "license", false, "show license")

	// Application specific options
	flag.BoolVar(&showTemplate, "s", false, "display the default template")
	flag.BoolVar(&showTemplate, "show-template", false, "display the default template")
	flag.StringVar(&templateFNames, "t", "", "colon delimited list of templates to use")
	flag.StringVar(&templateFNames, "templates", "", "colon delimited list of templates to use")
}

func main() {
	var (
		err error
	)

	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "MKPAGE", fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version), mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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
		fmt.Println(mkpage.DefaultTemplateSource)
		os.Exit(0)
	}

	var (
		// This holds the template sources
		templateSources []string
		// NOTE: Now we assemble everything with this template variable.
		tmpl *template.Template
	)

	// Make sure we have a configured command to run
	templateFNames = cfg.MergeEnv("templates", templateFNames)
	if len(templateFNames) > 0 {
		for _, fname := range strings.Split(templateFNames, ":") {
			templateSources = append(templateSources, fname)
		}
	}

	data := make(map[string]string)
	for i, arg := range args {
		if strings.Contains(arg, "=") == true {
			// Update data map
			pair := strings.SplitN(arg, "=", 2)
			if len(pair) != 2 {
				fmt.Fprintf(os.Stderr, "Can't read pair (%d) %s\n", i+1, arg)
				os.Exit(1)
			}
			data[pair[0]] = pair[1]
		} else {
			// Must be the template source
			templateSources = append(templateSources, arg)
		}
	}

	// Assemble the template(s)
	if len(templateSources) == 0 {
		tmpl, err = template.New("default.tmpl").Funcs(tmplfn.Join(tmplfn.TimeMap, tmplfn.PageMap)).Parse(mkpage.DefaultTemplateSource)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Template parsing failed, %s\n", err)
			os.Exit(1)
		}
	} else {
		//NOTE: The first template gets to hold the template functions.
		tmpl, err = template.New(templateSources[0]).Funcs(tmplfn.Join(tmplfn.TimeMap, tmplfn.PageMap)).ParseFiles(templateSources...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Template parsing failed, %s\n", err)
			os.Exit(1)
		}
	}
	// Make the page
	if err := mkpage.MakePage(os.Stdout, tmpl, data); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
