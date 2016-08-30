#
# Simple Makefile
#
build:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	gofmt -w cmds/reldocpath/reldocpath.go
	go build
	go build -o bin/mkpage cmds/mkpage/mkpage.go
	go build -o bin/reldocpath cmds/reldocpath/reldocpath.go

test:
	go test

save:
	./mk-website.bash
	git commit -am "quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f mkpage-binary-release.zip ]; then rm -f mkpage-binary-release.zip; fi

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
