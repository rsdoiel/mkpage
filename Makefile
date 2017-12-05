#
# Simple Makefile
#

PROJECT = mkpage

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

PKGASSETS = $(shell which pkgassets)

OS = $(shell uname)

EXT =
ifeq ($(OS),Windows)
	EXT = .exe
endif

build: bin/mkpage$(EXT) bin/mkslides$(EXT) bin/mkrss$(EXT) \
	bin/sitemapper$(EXT) bin/byline$(EXT) bin/titleline$(EXT) \
	bin/reldocpath$(EXT) bin/urlencode$(EXT) bin/urldecode$(EXT) \
	bin/ws$(EXT) 

mkpage.go: assets.go codesnip.go

assets.go:
	pkgassets -o assets.go -p mkpage Defaults defaults
	git add assets.go

bin/mkpage$(EXT): mkpage.go assets.go codesnip.go cmds/mkpage/mkpage.go
	go build -o bin/mkpage$(EXT) cmds/mkpage/mkpage.go

bin/mkslides$(EXT): mkpage.go cmds/mkslides/mkslides.go
	go build -o bin/mkslides$(EXT) cmds/mkslides/mkslides.go

bin/mkrss$(EXT): mkpage.go cmds/mkrss/mkrss.go
	go build -o bin/mkrss$(EXT) cmds/mkrss/mkrss.go

bin/sitemapper$(EXT): mkpage.go cmds/sitemapper/sitemapper.go
	go build -o bin/sitemapper$(EXT) cmds/sitemapper/sitemapper.go

bin/byline$(EXT): mkpage.go cmds/byline/byline.go
	go build -o bin/byline$(EXT) cmds/byline/byline.go

bin/titleline$(EXT): mkpage.go cmds/titleline/titleline.go
	go build -o bin/titleline$(EXT) cmds/titleline/titleline.go

bin/reldocpath$(EXT): cmds/reldocpath/reldocpath.go
	go build -o bin/reldocpath$(EXT) cmds/reldocpath/reldocpath.go

bin/urlencode$(EXT): cmds/urlencode/urlencode.go
	go build -o bin/urlencode$(EXT) cmds/urlencode/urlencode.go

bin/urldecode$(EXT): cmds/urldecode/urldecode.go
	go build -o bin/urldecode$(EXT) cmds/urldecode/urldecode.go

bin/ws$(EXT): mkpage.go cmds/ws/ws.go
	go build -o bin/ws$(EXT) cmds/ws/ws.go


lint:
	golint mkpage.go
	golint mkpage_test.go
	golint cmds/mkpage/mkpage.go
	golint cmds/mkslides/mkslides.go
	golint cmds/mkrss/mkrss.go
	golint cmds/sitemapper/sitemapper.go
	golint cmds/byline/byline.go
	golint cmds/titleline/titleline.go
	golint cmds/reldocpath/reldocpath.go
	golint cmds/urlencode/urlencode.go
	golint cmds/urldecode/urldecode.go
	golint cmds/ws/ws.go

format:
	gofmt -w mkpage.go
	gofmt -w mkpage_test.go
	gofmt -w cmds/mkpage/mkpage.go
	gofmt -w cmds/mkslides/mkslides.go
	gofmt -w cmds/mkrss/mkrss.go
	gofmt -w cmds/sitemapper/sitemapper.go
	gofmt -w cmds/byline/byline.go
	gofmt -w cmds/titleline/titleline.go
	gofmt -w cmds/reldocpath/reldocpath.go
	gofmt -w cmds/urlencode/urlencode.go
	gofmt -w cmds/urldecode/urldecode.go
	gofmt -w cmds/ws/ws.go

test: bin/mkpage$(EXT) bin/mkslides$(EXT) bin/mkrss$(EXT) \
	bin/sitemapper$(EXT) bin/byline$(EXT) bin/titleline$(EXT) \
	bin/reldocpath$(EXT) bin/urlencode$(EXT) bin/urldecode$(EXT) \
	bin/ws$(EXT) 
	go test
	bash test_cmds.bash

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ "$(PKGASSETS)" != "" ]; then rm assets.go; pkgassets -o assets.go -p mkpage Defaults defaults; git add assets.go; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(GOPATH)/bin go install cmds/mkpage/mkpage.go
	env GOBIN=$(GOPATH)/bin go install cmds/mkslides/mkslides.go
	env GOBIN=$(GOPATH)/bin go install cmds/mkrss/mkrss.go
	env GOBIN=$(GOPATH)/bin go install cmds/sitemapper/sitemapper.go
	env GOBIN=$(GOPATH)/bin go install cmds/byline/byline.go
	env GOBIN=$(GOPATH)/bin go install cmds/titleline/titleline.go
	env GOBIN=$(GOPATH)/bin go install cmds/reldocpath/reldocpath.go
	env GOBIN=$(GOPATH)/bin go install cmds/urlencode/urlencode.go
	env GOBIN=$(GOPATH)/bin go install cmds/urldecode/urldecode.go
	env GOBIN=$(GOPATH)/bin go install cmds/ws/ws.go


dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mkpage cmds/mkpage/mkpage.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mkslides cmds/mkslides/mkslides.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/mkrss cmds/mkrss/mkrss.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/sitemapper cmds/sitemapper/sitemapper.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/byline cmds/byline/byline.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/titleline cmds/titleline/titleline.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/reldocpath cmds/reldocpath/reldocpath.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/urlencode cmds/urlencode/urlencode.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/urldecode cmds/urldecode/urldecode.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/ws cmds/ws/ws.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* templates/*
	rm -fR dist/bin



dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mkpage.exe cmds/mkpage/mkpage.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mkslides.exe cmds/mkslides/mkslides.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/mkrss.exe cmds/mkrss/mkrss.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/sitemapper.exe cmds/sitemapper/sitemapper.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/byline.exe cmds/byline/byline.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/titleline.exe cmds/titleline/titleline.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/reldocpath.exe cmds/reldocpath/reldocpath.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/urlencode.exe cmds/urlencode/urlencode.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/urldecode.exe cmds/urldecode/urldecode.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/ws.exe cmds/ws/ws.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* templates/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mkpage cmds/mkpage/mkpage.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mkslides cmds/mkslides/mkslides.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/mkrss cmds/mkrss/mkrss.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/sitemapper cmds/sitemapper/sitemapper.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/byline cmds/byline/byline.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/titleline cmds/titleline/titleline.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/reldocpath cmds/reldocpath/reldocpath.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/urlencode cmds/urlencode/urlencode.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/urldecode cmds/urldecode/urldecode.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/ws cmds/ws/ws.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* templates/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mkpage cmds/mkpage/mkpage.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mkslides cmds/mkslides/mkslides.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/mkrss cmds/mkrss/mkrss.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/sitemapper cmds/sitemapper/sitemapper.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/byline cmds/byline/byline.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/titleline cmds/titleline/titleline.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/reldocpath cmds/reldocpath/reldocpath.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/urlencode cmds/urlencode/urlencode.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/urldecode cmds/urldecode/urldecode.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/ws cmds/ws/ws.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/* docs/* how-to/* templates/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist/docs
	mkdir -p dist/how-to
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v docs/byline.md dist/docs/
	cp -v docs/mkpage.md dist/docs/
	cp -v docs/mkrss.md dist/docs
	cp -v docs/mkslides.md dist/docs/
	cp -v docs/reldocpath.md dist/docs/
	cp -v docs/sitemapper.md dist/docs/
	cp -v docs/titleline.md dist/docs/
	cp -v docs/urldecode.md dist/docs/
	cp -v docs/urlencode.md dist/docs/
	cp -v docs/ws.md dist/docs/
	cp -v how-to/go-template-recipes.md dist/how-to/
	cp -v how-to/the-basics.md dist/how-to/
	cp -vR templates dist/
	./package-versions.bash > dist/package-versions.txt

release: assets.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

