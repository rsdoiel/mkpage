
# USAGE

    ws [OPTIONS] [DOCROOT]

## SYNOPSIS

	a nimble web server

ws is a command line utility for developing and testing static websites.
It uses Go's standard http libraries and can supports both http 1 and 2
out of the box.  It is intended as a minimal wrapper for Go's standard
http libraries supporting http/https versions 1 and 2 out of the box.

## CONFIGURATION

ws can be configurated through environment settings. The following are
supported.

+ MKPAGE_URL  - sets the URL to listen on (e.g. http://localhost:8000)
+ MKPAGE_DOCROOT - sets the document path to use
+ MKPAGE_SSL_KEY - the path to the SSL key if using https
+ MKPAGE_SSL_CERT - the path to the SSL cert if using https

## OPTIONS

```
	-c	Set the path for the SSL Cert
	-cert	Set the path for the SSL Cert
	-d	Set the htdocs path
	-docs	Set the htdocs path
	-h	Display this help message
	-help	Display this help message
	-k	Set the path for the SSL Key
	-key	Set the path for the SSL Key
	-l	Should license info
	-license	Should license info
	-u	The protocal and hostname listen for as a URL
	-url	The protocal and hostname listen for as a URL
	-v	Should version info
	-version	Should version info
```

## EXAMPLES

Run web server using the content in the current directory
(assumes the environment variables MKPAGE_DOCROOT are not defined).

```shell
   ws
```

Run web server using a specified directory

```shell
   ws /www/htdocs
```

