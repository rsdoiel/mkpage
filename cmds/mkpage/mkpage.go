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

	// My package
	"github.com/caltechlibrary/mkpage"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
)

func usage(fp *os.File, appName string) {
	fmt.Fprintf(fp, `
 USAGE: %s [OPTION] [KEY/VALUE DATA PAIRS] TEMPLATE_FILENAME [TEMPLATE_FILENAMES]

 Using the key value pairs populate the template(s) and render to stdout.

 OPTIONS

`, appName)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("    -%s %s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fp, `

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

	mkpage "now=text:$(date)" "name=text:Little Frieda" \
		"weather=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" \
		signature=testdata/signature.txt \
		testdata/myformletter.template

 Golang's text/template docs can be found at 

      https://golang.org/pkg/text/template/

 Version %s

`, mkpage.Version)
}

func license(fp *os.File, appName string) {
	fmt.Fprintf(fp, `
%s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`, appName)
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showLicense, "l", false, "show license")
}

func main() {
	var (
		err error
	)

	appName := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(os.Stdout, appName)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf(" Version %s\n", mkpage.Version)
		os.Exit(0)
	}
	if showLicense == true {
		license(os.Stdout, appName)
		os.Exit(0)
	}

	var templateSources []string
	data := make(map[string]string)
	args := flag.Args()
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
	if len(templateSources) == 0 {
		usage(os.Stderr, appName)
		fmt.Fprintln(os.Stderr, "ERROR: Missing a page template")
		os.Exit(1)
	}

	tmpl, err := template.ParseFiles(templateSources...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Template parsing failed, %s\n", err)
		os.Exit(1)
	}
	if err := mkpage.MakePage(os.Stdout, tmpl, data); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
