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
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	// My packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/mkpage"
	"github.com/rsdoiel/rss2"
)

var (
	// Standard options
	showHelp    bool
	showLicense bool
	showVersion bool

	// App specific options
	excludeList        string
	articleLimit       int
	channelLanguage    string
	channelTitle       string
	channelDescription string
	channelLink        string
	channelGenerator   string
	channelPubDate     string
	channelBuildDate   string
	channelCopyright   string
	channelCategory    string

	// Usage and docs
	usage = `USAGE: %s [OPTION] HTDOCS [RSS_FILENAME]`

	description = `
SYNOPSIS

%s walks the file system to generate a RSS2 file. It assumes that the directory
for HTDOCS is is the base directory containing subdirectories in the form of
/YYYY/MM/DD/ARTICLE_HTML where YYYY/MM/DD (Year, Month, Day) 
corresponds to the publication date of ARTICLE_HTML.
`

	examples = `
EXAMPLE

If our htdocs folder is our document root and out blog is 
htdocs/myblog.

    %s -channel-title="This Great Beyond" \
	   -channel-description="Blog to save the world" \
	   -channel-link="http://blog.example.org" \
	   htdocs htdocs/rss.xml 

This would build an RSS 2 file in htdocs/rss.xml from the 
articles in htdocs/myblog/YYYY/MM/DD.
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
	flag.StringVar(&channelLanguage, "channel-language", "", "Language, e.g. en-ca")
	flag.StringVar(&channelTitle, "channel-title", "", "Title of channel")
	flag.StringVar(&channelDescription, "channel-description", "", "Description of channel")
	flag.StringVar(&channelLink, "channel-link", "", "link to channel")
	flag.StringVar(&channelGenerator, "channel-generator", "", "Name of RSS generator")
	flag.StringVar(&channelPubDate, "channel-pubdate", "", "Pub Date for channel")
	flag.StringVar(&channelBuildDate, "channel-builddate", "", "Build Date for channel")
	flag.StringVar(&channelCopyright, "channel-copyright", "", "Copyright for channel")
	flag.StringVar(&channelCategory, "channel-category", "", "category for channel")
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

	if len(channelTitle) == 0 {
		channelTitle = `A website`
	}
	if len(channelDescription) == 0 {
		channelDescription = `A collection of web pages`
	}
	if len(channelLink) == 0 {
		channelLink = `http://localhost:8000`
	}

	// Setup the Channel metadata for feed.
	feed := new(rss2.RSS2)
	feed.Version = "2.0"
	feed.Title = channelTitle
	feed.Description = channelDescription
	feed.Link = channelLink
	if len(channelLanguage) > 0 {
		feed.Language = channelLanguage
	}
	if len(channelCopyright) > 0 {
		feed.Copyright = channelCopyright
	}
	if len(channelCategory) > 0 {
		feed.Category = channelCategory
	}
	if len(channelGenerator) == 0 {
		feed.Generator = cfg.Version()
	} else {
		feed.Generator = channelGenerator
	}
	now := time.Now()
	if len(channelPubDate) == 0 {
		feed.PubDate = now.Format(time.RFC822)
	} else {
		feed.PubDate = channelPubDate
	}
	if len(channelBuildDate) == 0 {
		feed.LastBuildDate = now.Format(time.RFC822)
	} else {
		feed.LastBuildDate = channelBuildDate
	}

	// Process command line parameters
	htdocs := "."
	rssPath := ""
	if len(args) > 0 {
		htdocs = args[0]
	}
	if len(args) > 1 {
		rssPath = args[1]
	}

	validBlogPath := regexp.MustCompile("/[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]/")
	err := mkpage.Walk(htdocs, func(p string, info os.FileInfo) bool {
		fname := path.Base(p)
		if validBlogPath.MatchString(p) == true &&
			strings.HasSuffix(fname, ".html") == true &&
			strings.HasPrefix(fname, "40") == false &&
			strings.HasPrefix(fname, "50") == false {
			return true
		}
		return false
	}, func(p string, info os.FileInfo) error {
		pname := strings.TrimPrefix(p, htdocs)
		if strings.HasPrefix(pname, "/") {
			pname = strings.TrimPrefix(pname, "/")
		}
		articleURL := fmt.Sprintf("%s/%s", channelLink, pname)
		u, err := url.Parse(articleURL)
		if err != nil {
			return err
		}
		title := mkpage.Unslugify(strings.TrimSuffix(path.Base(u.Path), ".html"))
		item := new(rss2.Item)
		item.Title = title
		item.Link = articleURL
		//FIXME: Need to pull the description from the Markdown version
		feed.ItemList = append(feed.ItemList, *item)
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Marshal RSS2 and render output
	src, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if len(rssPath) > 0 {
		if err := ioutil.WriteFile(rssPath, src, 0664); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	fmt.Fprintf(os.Stdout, "%s\n", src)
}
