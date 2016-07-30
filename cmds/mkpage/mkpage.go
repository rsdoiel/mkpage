package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// My package
	"github.com/rsdoiel/mkpage"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	useHTMLTemplate      bool
	useMarkdownProcessor bool
)

func usage(fp *os.File, appName string) {
	fmt.Fprintf(fp, `
 USAGE: %s [OPTION] [KEY/VALUE DATA PAIRS] TEMPLATE_FILENAME

 Using the key value pairs populate the template and render to stdout.

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

	mkpage "now=string:$(date)" "name=string:Little Frieda" \
		"weather=http://forecast.weather.gov/MapClick.php?lat=9.9667&lon=139.6667&FcstType=json" \
		signature=test/signature.txt \
		testdata/myformletter.template

 Golang's template docs can be found at 

 + https://golang.org/pkg/text/template/
 + https://golang.org/pkg/html/template/

 Version %s

`, mkpage.Version)
}

func license(fp *os.File, appName string) {
	fmt.Fprintf(fp, `
%s

Copyright (c) 2016, R. S. Doiel
All rights not granted herein are expressly reserved by R. S. Doiel.

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

	flag.BoolVar(&useHTMLTemplate, "html", false, "use Go's html/template instead of text/template")
	flag.BoolVar(&useMarkdownProcessor, "m", false, "apply markdown processor to \".md\" files")
	flag.BoolVar(&useMarkdownProcessor, "markdown", false, "apply markdown processor to \".md\" files")
}

func main() {
	var (
		src []byte
		err error
	)

	appName := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(os.Stdout, appName)
	}
	if showVersion == true {
		fmt.Printf(" Version %s\n", mkpage.Version)
	}
	if showLicense == true {
		license(os.Stdout, appName)
	}

	templateType := mkpage.Text
	if useHTMLTemplate == true {
		templateType = mkpage.HTML
	}
	data := make(map[string][]byte)
	args := flag.Args()
	for i, arg := range args {
		if strings.Contains(arg, "=") == true {
			// Update data map
			pair := strings.SplitN(arg, "=", 2)
			if len(pair) != 2 {
				fmt.Fprintf(os.Stderr, "Can't read pair (%d) %s\n", i+1, arg)
				os.Exit(1)
			}
			data[pair[0]] = []byte(pair[1])
		} else {
			// Must be the template source
			src, err = ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't read %s %s", arg, err)
				os.Exit(1)
			}
		}
	}
	mkpage.MakePage(os.Stdout, string(src), templateType, data, useMarkdownProcessor)
}
