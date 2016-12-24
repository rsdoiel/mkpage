#
# Simple Makefile
#

PROJECT = mkpage

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\  -f 2)

build:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	gofmt -w cmds/reldocpath/reldocpath.go
	go build -o bin/mkpage cmds/mkpage/mkpage.go
	go build -o bin/reldocpath cmds/reldocpath/reldocpath.go

test:
	go test

save:
	git commit -am "quick save"
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/mkpage/mkpage.go
	env GOBIN=$(HOME)/bin go install cmds/reldocpath/reldocpath.go

release:
	./mk-release.bash

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash
