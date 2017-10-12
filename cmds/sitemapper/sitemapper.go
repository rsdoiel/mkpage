//
// sitemapper generates a sitemap.xml file by crawling the content generate with genpages
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
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkpage"
)

type locInfo struct {
	Loc     string
	LastMod string
}

var (
	usage = `USAGE: %s [OPTIONS] HTDOCS_PATH MAP_FILENAME PUBLIC_BASE_URL`

	description = `

SYNOPSIS

%s generates a sitemap for the website.

`

	examples = `

EXAMPLE

    %s htdocs htdocs/sitemap.xml http://eprints.example.edu

`

	// Standard options
	showHelp     bool
	showVersion  bool
	showLicense  bool
	showExamples bool

	// App options
	htdocs       string
	siteURL      string
	excludeList  string
	sitemapFName string

	changefreq string
	locList    []*locInfo
)

func check(cfg *cli.Config, key, value string) string {
	if value == "" {
		log.Fatalf("Missing %s_%s", cfg.EnvPrefix, strings.ToUpper(key))
		return ""
	}
	return value
}

// ExcludeList is a list of directories to skip when generating a sitemap
type ExcludeList []string

// Set returns the len of the new DirList array based on spliting the passed in string
func (dirList ExcludeList) Set(s string) int {
	if len(strings.TrimSpace(s)) > 0 {
		dirList = strings.Split(s, ":")
	}
	return len(dirList)
}

// Exclude returns true if a fname fragment is included in set of dirList
func (dirList ExcludeList) Exclude(p string) bool {
	for _, item := range dirList {
		if len(item) > 0 && len(p) > 0 && strings.Contains(p, item) == true {
			log.Printf("Skipping %q", p)
			return true
		}
	}
	return false
}

func init() {
	// Log to standard out
	log.SetOutput(os.Stdout)

	// Setup options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// App specific options
	flag.StringVar(&changefreq, "u", "daily", "Set the change frequencely value, e.g. daily, weekly, monthly")
	flag.StringVar(&changefreq, "update-frequency", "daily", "Set the change frequencely value, e.g. daily, weekly, monthly")
	flag.StringVar(&excludeList, "e", "", "A colon delimited list of path parts to exclude from sitemap")
	flag.StringVar(&excludeList, "exclude", "", "A colon delimited list of path parts to exclude from sitemap")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "MKPAGE", mkpage.Version)
	cfg.LicenseText = fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	if len(args) != 3 {
		fmt.Printf("%s requires 3 parameters, see %s --help\n", appName, appName)
		os.Exit(1)
	}

	if len(args) > 0 {
		htdocs = args[0]
	}
	if len(args) > 1 {
		sitemapFName = args[1]
	}
	if len(args) > 2 {
		siteURL = args[2]
	}

	// Required
	htdocs = check(cfg, "docroot", cfg.MergeEnv("docroot", htdocs))
	siteURL = check(cfg, "site_url", cfg.MergeEnv("site_url", siteURL))
	sitemapFName = check(cfg, "sitemap", cfg.MergeEnv("sitemap", sitemapFName))

	if changefreq == "" {
		changefreq = "daily"
	}

	site, err := url.Parse(siteURL)
	if err != nil {
		fmt.Printf("Invalid site URL: %q, %s\n", siteURL, err)
		os.Exit(1)
	}

	excludeDirs := ExcludeList(strings.Split(excludeList, ":"))

	log.Printf("Starting map of %s\n", htdocs)
	filepath.Walk(htdocs, func(p string, info os.FileInfo, err error) error {
		if strings.HasSuffix(p, ".html") {
			fname := path.Base(p)
			//NOTE: You can skip the eror pages, and excluded directories in the sitemap
			if strings.HasPrefix(fname, "50") == false && strings.HasPrefix(p, "40") == false && excludeDirs.Exclude(p) == false {
				finfo := new(locInfo)
				//FIXME: should use the parsed URL and append to path
				page, _ := url.Parse(site.String())
				page.Path = path.Join(page.Path, strings.TrimPrefix(p, htdocs))
				finfo.Loc = page.String()
				yr, mn, dy := info.ModTime().Date()
				finfo.LastMod = fmt.Sprintf("%d-%0.2d-%0.2d", yr, mn, dy)
				log.Printf("Adding %s\n", finfo.Loc)
				locList = append(locList, finfo)
			}
		}
		return nil
	})
	fmt.Printf("Writing %s\n", sitemapFName)
	fp, err := os.OpenFile(sitemapFName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatalf("Can't create %s, %s\n", sitemapFName, err)
	}
	defer fp.Close()
	fp.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
`))
	for _, item := range locList {
		fp.WriteString(fmt.Sprintf(`
    <url>
            <loc>%s</loc>
            <lastmod>%s</lastmod>
            <changefreq>%s</changefreq>
    </url>
`, item.Loc, item.LastMod, changefreq))
	}
	fp.Write([]byte(`
</urlset>
`))
}
