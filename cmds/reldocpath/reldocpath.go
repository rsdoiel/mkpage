//
// reldocpath.go takes a source document path and a target document path with same base path
// returning a relative path to the target file.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights not granted herein are expressly reserved by R. S. Doiel.
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
	"io"
	"os"
	"path"

	"github.com/rsdoiel/mkpage"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
)

func usage(out io.Writer, appname, version string) {
	fmt.Fprintf(out, `
 USAGE: %s SOURCE_DOC_PATH TARGET_DOC_PATH 

 Given a source document path, a target document path calculate and
 the implied common base path calculate the relative path for target.

 EXAMPLE:

 Given

     %s chapter-01/lesson-03.html css/site.css

 would output

     .../css/site.css

 OPTIONS

`, appname, appname)
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			fmt.Fprintf(out, "    -%s, --%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
		}
	})
	fmt.Fprintf(out, "\n\n Version %s\n", version)
}

func license(out io.Writer, appname, version string) {
	fmt.Fprintf(out, `
%s %s

Copyright (c) 2016, R. S. Doiel
All rights not granted herein are expressly reserved by R. S. Doiel.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`, appname, version)
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(os.Stdout, appname, mkpage.Version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Fprintf(os.Stdout, "%s %s\n", appname, mkpage.Version)
		os.Exit(0)
	}
	if showLicense == true {
		license(os.Stdout, appname, mkpage.Version)
		os.Exit(0)
	}
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, " Expected a source and target file path\n For help try: %s -help", appname)
		os.Exit(1)

	}
	source, target := args[0], args[1]
	fmt.Fprintf(os.Stdout, "%s", mkpage.RelativeDocPath(source, target))
}
