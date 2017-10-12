//
// ws.go - A simple web server for static files and limit server side JavaScript
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech
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
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/mkpage"

	// Other packages
	"golang.org/x/crypto/acme/autocert"
)

// Flag options
var (
	usage = `USAGE: %s [OPTIONS] [DOCROOT]`

	description = `

SYNOPSIS

	a nimble web server

%s is a command line utility for developing and testing static websites.
It uses Go's standard http libraries and can supports both http 1 and 2
out of the box.  It is intended as a minimal wrapper for Go's standard
http libraries supporting http/https versions 1 and 2 out of the box.

CONFIGURATION

%s can be configurated through environment settings. The following are
supported.

+ MKPAGE_URL  - sets the URL to listen on (e.g. http://localhost:8000)
+ MKPAGE_DOCROOT - sets the document path to use
+ MKPAGE_SSL_KEY - the path to the SSL key if using https
+ MKPAGE_SSL_CERT - the path to the SSL cert if using https

`

	examples = `

EXAMPLES

Run web server using the content in the current directory
(assumes the environment variables MKPAGE_DOCROOT are not defined).

   %s

Run web server using a specified directory

   %s /www/htdocs

Running web server using ACME TLS support (i.e. Let's Encrypt).
Note will only include the hostname as the ACME setup is for
listenning on port 443. This may require privilaged account
and will require that the hostname listed matches the public
DNS for the machine (this is need by the ACME protocol to
issue the cert, see https://letsencrypt.org for details)

   %s -acme -url www.example.org /www/htdocs

`

	// Standard options
	showHelp     bool
	showVersion  bool
	showLicense  bool
	showExamples bool

	// local app options
	uri         string
	docRoot     string
	sslKey      string
	sslCert     string
	letsEncrypt bool
)

func logRequest(r *http.Request) {
	log.Printf("Request: %s Path: %s RemoteAddr: %s UserAgent: %s\n", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		next.ServeHTTP(w, r)
	})
}

func init() {
	defaultDocRoot := "."
	defaultURL := "http://localhost:8000"

	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// Application Options
	flag.StringVar(&docRoot, "d", defaultDocRoot, "Set the htdocs path")
	flag.StringVar(&docRoot, "docs", defaultDocRoot, "Set the htdocs path")
	flag.StringVar(&uri, "u", defaultURL, "The protocal and hostname listen for as a URL")
	flag.StringVar(&uri, "url", defaultURL, "The protocal and hostname listen for as a URL")
	flag.StringVar(&sslKey, "k", "", "Set the path for the SSL Key")
	flag.StringVar(&sslKey, "key", "", "Set the path for the SSL Key")
	flag.StringVar(&sslCert, "c", "", "Set the path for the SSL Cert")
	flag.StringVar(&sslCert, "cert", "", "Set the path for the SSL Cert")
	flag.BoolVar(&letsEncrypt, "acme", false, "Enable Let's Encypt ACME TLS support")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "MKPAGE", mkpage.Version)
	cfg.LicenseText = fmt.Sprintf(mkpage.LicenseText, appName, mkpage.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName)
	cfg.OptionText = "OPTIONS"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName)

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

	// setup from command line
	if len(args) > 0 {
		docRoot = args[0]
	}

	docRoot = cfg.CheckOption("dotroot", cfg.MergeEnv("docroot", docRoot), true)
	log.Printf("DocRoot %s", docRoot)

	uri = cfg.CheckOption("url", cfg.MergeEnv("url", uri), true)
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Can't parse %q, %s", uri, err)
	}

	if u.Scheme == "https" && letsEncrypt == false {
		sslKey = cfg.CheckOption("ssl_key", cfg.MergeEnv("ssl_key", sslKey), true)
		sslCert = cfg.CheckOption("ssl_cert", cfg.MergeEnv("ssl_cert", sslCert), true)
		log.Printf("SSL Key %s", sslKey)
		log.Printf("SSL Cert %s", sslCert)
	}
	log.Printf("Listening for %s", uri)

	http.Handle("/", http.FileServer(http.Dir(docRoot)))
	if letsEncrypt == true {
		// Note: use a sensible value for data directory
		// this is where cached certificates are stored
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Can't determine current working directory before creating etc/acme")
		}
		if docRoot == "." || docRoot == cwd {
			log.Fatal("Can't create etc/acme in shared document root")
		}
		cacheDir := "etc/acme"
		os.MkdirAll(cacheDir, 0700)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(u.Host),
			Cache:      autocert.DirCache(cacheDir),
		}
		sSvr := &http.Server{
			Addr:      ":https",
			TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
			Handler:   mkpage.RequestLogger(mkpage.StaticRouter(http.DefaultServeMux)),
		}
		// Launch the TLS version
		go func() {
			log.Fatal(sSvr.ListenAndServeTLS("", ""))
		}()

		rmux := http.NewServeMux()
		rmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var target string
			if strings.HasPrefix(r.URL.Path, "/") == false {
				target = u.String() + "/" + r.URL.Path
			} else {
				target = u.String() + r.URL.Path
			}
			if len(r.URL.RawQuery) > 0 {
				target += "?" + r.URL.RawQuery
			}
			mkpage.ResponseLogger(r, http.StatusTemporaryRedirect, fmt.Errorf("redirecting %s to %s", r.URL.String(), target))
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
		})
		pSvr := &http.Server{
			Addr:    ":http",
			Handler: mkpage.RequestLogger(rmux),
		}
		log.Printf("Redirecting http://%s to to %s", u.Host, u.String())
		log.Fatal(pSvr.ListenAndServe())
	} else if u.Scheme == "https" {
		err := http.ListenAndServeTLS(u.Host, sslCert, sslKey, mkpage.RequestLogger(mkpage.StaticRouter(http.DefaultServeMux)))
		if err != nil {
			log.Fatalf("%s", err)
		}
	} else {
		err := http.ListenAndServe(u.Host, mkpage.RequestLogger(mkpage.StaticRouter(http.DefaultServeMux)))
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
}
