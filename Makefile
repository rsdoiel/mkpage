#
# Simple Makefile
#

PROJECT = mkpage

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build:
	go build -o bin/mkpage cmds/mkpage/mkpage.go
	go build -o bin/reldocpath cmds/reldocpath/reldocpath.go
	go build -o bin/slugify cmds/slugify/slugify.go

lint:
	golint mkpage.go
	golint mkpage_test.go
	golint cmds/mkpage/mkpage.go
	golint cmds/reldocpath/reldocpath.go
	golint cmds/slugify/slugify.go

format:
	goimports -w mkpage.go
	goimports -w mkpage_test.go
	goimports -w cmds/mkpage/mkpage.go
	goimports -w cmds/reldocpath/reldocpath.go
	goimports -w cmds/slugify/slugify.go
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	gofmt -w cmds/reldocpath/reldocpath.go
	gofmt -w cmds/slugify/slugify.go

test:
	go test

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/mkpage/mkpage.go
	env GOBIN=$(HOME)/bin go install cmds/reldocpath/reldocpath.go
	env GOBIN=$(HOME)/bin go install cmds/slugify/slugify.go


dist/linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/slugify cmds/slugify/slugify.go

dist/windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/mkpage.exe cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/reldocpath.exe cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/slugify.exe cmds/slugify/slugify.go

dist/macosx-adm64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/slugify cmds/slugify/slugify.go

dist/raspbian-arm7:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/slugify cmds/slugify/slugify.go

dist/raspbian-arm6:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/slugify cmds/slugify/slugify.go

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/raspbian-arm6
	mkdir -p dist/etc
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v etc/*-example dist/etc/
	cp -vR scripts dist/
	cp -vR templates dist/
	cp -vR examples dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

