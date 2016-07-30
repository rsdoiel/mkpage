#
# Simple Makefile
#
build:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	go build
	go build -o bin/mkpage cmds/mkpage/mkpage.go

test:
	go test

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f mkpage-binary-release.zip ]; then rm -f mkpage-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/mkpage/mkpage.go

release:
	./mk-release.bash

publish:
	./mk-website.bash
	./publish.bash
