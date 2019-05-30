@echo off
REM This is a Windows 10 Batch file for building mkpage tools
REM from the command prompt.
REM
REM It requires: go version 1.12.4 or better and the cli for git installed
REM
go version
echo "Getting ready to build the commands and write them to .\bin"
@echo on

go build -o bin\byline.exe cmd\byline\byline.go
go build -o bin\mkpage.exe cmd\mkpage\mkpage.go
go build -o bin\mkrss.exe cmd\mkrss\mkrss.go
go build -o bin\mkslides.exe cmd\mkslides\mkslides.go
go build -o bin\reldocpath.exe cmd\reldocpath\reldocpath.go
go build -o bin\sitemapper.exe cmd\sitemapper\sitemapper.go
go build -o bin\titleline.exe cmd\titleline\titleline.go
go build -o bin\urldecode.exe cmd\urldecode\urldecode.go
go build -o bin\urlencode.exe cmd\urlencode\urlencode.go
go build -o bin\ws.exe cmd\ws\ws.go

@echo off
echo "You can now copy the contents of .\bin to your program directory"
