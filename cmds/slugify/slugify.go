//
// slugify.go turn a human readable string into a URL/path friendly string
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
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
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkpage"
)

var (
	usage = `USAGE: %s [OPTIONS] STRING_TO_SLUGIFY`

	description = `
SYNOPSIS

% changes a human readable string into a path or URL friendly
string. E.g. "Hello World" becomes "hello-world"
`

	examples = `
EXAMPLE

    %s "Hello World my friend"
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// Application Options
	mixedCase bool
	overrides string

	// substitutions
	substitutions = map[string]string{
		" ": "-",
		"/": "_",
		"'": "",
		`"`: "",
	}
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")

	// Application Options
	flag.StringVar(&overrides, "s", "", "JSON representation of substitution")
	flag.StringVar(&overrides, "substitutions", "", "JSON representation of substitutions")
	flag.BoolVar(&mixedCase, "m", true, "allow mixed case, defats to true")
	flag.BoolVar(&mixedCase, "mixed-case", true, "allow mixed case, defaults to true")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	// Configuration and command line interation
	cfg := cli.New(appName, "MKPAGE", fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version), mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	args := flag.Args()
	phrase := strings.Join(args, " ")

	if len(overrides) > 0 {
		opt := map[string]string{}
		err := json.Unmarshal([]byte(overrides), &opt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		for k, v := range opt {
			substitutions[k] = v
		}
	}

	for target, replacement := range substitutions {
		phrase = strings.Replace(phrase, target, replacement, -1)
	}
	if mixedCase == false {
		phrase = strings.ToLower(phrase)
	}
	fmt.Println(phrase)
}
