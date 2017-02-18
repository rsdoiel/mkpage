#
# Simple Makefile
#

PROJECT = mkpage

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build: bin/mkpage bin/reldocpath bin/slugify bin/mkslides bin/sitemapper bin/mkrss bin/byline bin/titleline bin/ws

bin/mkpage: mkpage.go cmds/mkpage/mkpage.go
	go build -o bin/mkpage cmds/mkpage/mkpage.go

bin/reldocpath: cmds/reldocpath/reldocpath.go
	go build -o bin/reldocpath cmds/reldocpath/reldocpath.go

bin/slugify: cmds/slugify/slugify.go
	go build -o bin/slugify cmds/slugify/slugify.go

bin/mkslides: mkpage.go cmds/mkslides/mkslides.go
	go build -o bin/mkslides cmds/mkslides/mkslides.go

bin/sitemapper: mkpage.go cmds/sitemapper/sitemapper.go
	go build -o bin/sitemapper cmds/sitemapper/sitemapper.go

bin/mkrss: mkpage.go cmds/mkrss/mkrss.go
	go build -o bin/mkrss cmds/mkrss/mkrss.go

bin/byline: mkpage.go cmds/byline/byline.go
	go build -o bin/byline cmds/byline/byline.go

bin/titleline: mkpage.go cmds/titleline/titleline.go
	go build -o bin/titleline cmds/titleline/titleline.go

bin/ws: mkpage.go cmds/ws/ws.go
	go build -o bin/ws cmds/ws/ws.go

lint:
	golint mkpage.go
	golint mkpage_test.go
	golint cmds/mkpage/mkpage.go
	golint cmds/reldocpath/reldocpath.go
	golint cmds/slugify/slugify.go
	golint cmds/mkslides/mkslides.go
	golint cmds/sitemapper/sitemapper.go
	golint cmds/mkrss/mkrss.go
	golint cmds/byline/byline.go
	golint cmds/titleline/titleline.go
	golint cmds/ws/ws.go

format:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	gofmt -w cmds/reldocpath/reldocpath.go
	gofmt -w cmds/slugify/slugify.go
	gofmt -w cmds/mkslides/mkslides.go
	gofmt -w cmds/sitemapper/sitemapper.go
	gofmt -w cmds/mkrss/mkrss.go
	gofmt -w cmds/byline/byline.go
	gofmt -w cmds/titleline/titleline.go
	gofmt -w cmds/ws/ws.go

test:
	go test

status:
	git status

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
	env GOBIN=$(HOME)/bin go install cmds/mkslides/mkslides.go
	env GOBIN=$(HOME)/bin go install cmds/sitemapper/sitemapper.go
	env GOBIN=$(HOME)/bin go install cmds/mkrss/mkrss.go
	env GOBIN=$(HOME)/bin go install cmds/byline/byline.go
	env GOBIN=$(HOME)/bin go install cmds/titleline/titleline.go
	env GOBIN=$(HOME)/bin go install cmds/ws/ws.go


dist/linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/slugify cmds/slugify/slugify.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/mkslides cmds/mkslides/mkslides.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/sitemapper cmds/sitemapper/sitemapper.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/mkrss cmds/mkrss/mkrss.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/byline cmds/byline/byline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/titleline cmds/titleline/titleline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/ws cmds/ws/ws.go

dist/windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/mkpage.exe cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/reldocpath.exe cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/slugify.exe cmds/slugify/slugify.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/mkslides.exe cmds/mkslides/mkslides.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/sitemapper.exe cmds/sitemapper/sitemapper.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/mkrss.exe cmds/mkrss/mkrss.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/byline.exe cmds/byline/byline.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/titleline.exe cmds/titleline/titleline.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/ws.exe cmds/ws/ws.go

dist/macosx-amd64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/slugify cmds/slugify/slugify.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/mkslides cmds/mkslides/mkslides.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/sitemapper cmds/sitemapper/sitemapper.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/mkrss cmds/mkrss/mkrss.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/byline cmds/byline/byline.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/titleline cmds/titleline/titleline.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/ws cmds/ws/ws.go

dist/raspbian-arm7:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/slugify cmds/slugify/slugify.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/mkslides cmds/mkslides/mkslides.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/sitemapper cmds/sitemapper/sitemapper.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/mkrss cmds/mkrss/mkrss.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/byline cmds/byline/byline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/titleline cmds/titleline/titleline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/ws cmds/ws/ws.go

dist/raspbian-arm6:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/mkpage cmds/mkpage/mkpage.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/reldocpath cmds/reldocpath/reldocpath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/slugify cmds/slugify/slugify.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/mkslides cmds/mkslides/mkslides.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/sitemapper cmds/sitemapper/sitemapper.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/mkrss cmds/mkrss/mkrss.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/byline cmds/byline/byline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/titleline cmds/titleline/titleline.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/ws cmds/ws/ws.go

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/raspbian-arm6
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v mkslides.md dist/
	cp -v sitemapper.md dist/
	cp -v reldocpath.md dist/
	cp -v slugify.md dist/
	cp -v byline.md dist/
	cp -v titleline.md dist/
	cp -v ws.md dist/
	cp -vR demo dist/
	cp -vR examples dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

