#
# Simple Makefile
#
PROJECT = mkpage

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build:
	env CGO_ENABLED=0 go build -o bin/mkpage cmds/mkpage/mkpage.go

lint:
	golint mkpage.go
	golint cmds/mkpage/mkpage.go

test:
	go test

save:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	git commit -am "quick save"
	git push origin $(VERSION)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/mkpage/mkpage.go

release:
	./mk-release.bash

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash
