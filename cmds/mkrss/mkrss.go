//
// mkrss.go is a command line tool for generating an RSS file from a blog
// directory structure in the form of PATH_TO_BLOG/YYYY/MM/DD/BLOG_ARTICLES.html
//
// @Author R. S. Doiel, <rsdoiel@gmail.com>
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
	"os"
	"path"

	// My packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/mkpage"
)

var (
	// Standard options
	showHelp    bool
	showLicense bool
	showVersion bool

	// App specific options
	excludeList  string
	articleLimit int

	// Usage and docs
	usage = `USAGE: %s [OPTION] HTDOCS BLOG_PATH RSS_FILENAME BASE_URL`

	description = `
SYNOPSIS

%s walks the file system to generate a RSS2 file. It assumes that the directory
for BLOG_PATH is is the base directory conforming to 
BLOG_PATH/YYYY/MM/DD/ARTICLE_HTML where YYYY/MM/DD (Year, Month, Day) 
corresponds to the publication date of ARTICLE_HTML.
`

	examples = `
EXAMPLE

    %s htdocs htdocs/myblog htdocs/blog.rss http://blog.example.org

This would build an RSS 2 file in htdocs/blog.rss from the articles in
htdocs/myblog/YYYY/MM/DD.
`
)

func init() {
	// Standard options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")

	// App specific options
	flag.StringVar(&excludeList, "e", "", "A colon delimited list of path exclusions")
	flag.IntVar(&articleLimit, "c", 0, "If non-zero, limit the number of articles in the RSS file")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	cfg := cli.New(appName, appName, fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version), mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName)

	args := flag.Args()

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

	// Now get the command line args
	if len(args) != 4 {
		fmt.Fprintf(os.Stderr, "Expecting four parameters, try %s -help for details\n", appName)
		os.Exit(1)
	}

	// Required Params
	htdocs := args[0]
	blogPath := args[1]
	rssPath := args[2]
	siteURL := args[3]
	fmt.Printf("DEBUG %s, %s, %s, %s\n", htdocs, blogPath, rssPath, siteURL)

}
