//
// urlencode.go is a simple command line utility to encode
// a string in a URL friendly way.
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
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

	// My Packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/mkpage"
)

var (
	usage = `USAGE: %s [OPTIONS] [STRING_TO_ENCODE]`

	description = `
SYNOPSIS

%s is a simple command line utility to URL encode content. By default
it reads from standard input and writes to standard out.  You can
also specifty the string to encode as a command line parameter.
`

	examples = `
EXAMPLES

    echo "This is the string to encode & nothing else!" | %s

would yield

    This%%20is%%20the%%20string%%20to%%20encode%%20&%%20nothing%%20else%%0A
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string

	// App Options
	useQueryEscape bool
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&inputFName, "i", "", "set input filename")
	flag.StringVar(&inputFName, "input", "", "set input filename")
	flag.StringVar(&outputFName, "o", "", "set output filename")
	flag.StringVar(&outputFName, "output", "", "set output filename")

	// App Options
	flag.BoolVar(&useQueryEscape, "q", false, "use query escape (pluses for spaces)")
	flag.BoolVar(&useQueryEscape, "query", false, "use query escape (pluses for spaces)")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Populate cfg from the environment
	cfg := cli.New(appName, "MKPAGE", fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version), mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

	// Handle the default options
	if showHelp == true {
		fmt.Println(cfg.Usage())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	in, err := cli.Open(inputFName, os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(inputFName, in)

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprint(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	var src string

	if len(args) > 0 {
		src = strings.Join(args, " ")
	} else {
		buf, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Fprint(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		src = fmt.Sprintf("%s", buf)
	}
	if useQueryEscape == true {
		fmt.Fprintf(out, "%s", url.QueryEscape(src))
	} else {
		fmt.Fprintf(out, "%s", url.PathEscape(src))
	}
}
