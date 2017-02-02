
# ws

## A nimble web server

_ws_ is a prototyping platform for web based services and websites.

### _ws_ has a minimal feature set

+ A simple static file webserver 
    + quick startup
    + activity logged to the console
    + supports http2 out of the box
+ A project setup option called *init*


## Configuration

You can configure _ws_ with command line options or environment variables.
Try `ws -help` for a list of command line options.

### Environment variables

+ WS_URL the URL to listen for by _ws_
  + default is http://localhost:8000
+ WS_DOCROOT the directory of your static content you need to serve
  + the default is ./htdocs
+ WS_SSL_KEY the path the the SSL key file (e.g. etc/ssl/site.key)
  + default is empty, only checked if your WS_URL is starts with https://
+ WS_SSL_CERT the path the the SSL cert file (e.g. etc/ssl/site.crt)
  + default is empty, only checked if your WS_URL is starts with https://

### Command line options

+ -url overrides WS_URL
+ -docs overrides WS_DOCROOT
+ -ssl-key overrides WS_SSL_KEY
+ -ssl-pem overrides WS_SSL_PEM
+ -init triggers the initialization process and creates a setup.bash file
+ -h, -help displays the help documentation
+ -v, -version display the version of the command
+ -l, -license displays license information

Running _ws_ without environment variables or command line options is an easy way
to server your current working directory's content out as http://localhost:8000.

Need to quickly build out a website from Markdown files or other JSON resources?
Take a look at [mkpage](https://caltechlibrary.github.io/mkpage).


## Installation

_ws_ is available as precompile binaries for Linux, Mac OS X, and Windows 10 on Intel.
Additional binaries are provided for Raspbian on ARM6 adn ARM7.  Follow the [INSTALL.md](install.html) 
instructions to download and install the pre-compiled binaries.

If you have Golang installed then _ws_ can be installed with the *go get* command.

```
    go get github.com/caltechlibrary/ws/...
```

## Compiling from source

Required

+ [Golang](http://golang.org) version 1.7 or better

Here's my basic approach to get things setup. Go 1.7 needs to be already installed.

```
  git clone https://github.com/caltechlibrary/ws
  cd ws
  go test
  go build
  go build cmds/ws/ws.go
```

If everything compiles fine then I do something like this--

```
  go install cmds/ws/ws.go
```


## LICENSE

copyright (c) 2016 Caltech
See [LICENSE](license.html) for details

