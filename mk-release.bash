#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM6 and Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
RELEASE_NAME=mkpage
echo "NOTE: this can take a while..."
for PROGNAME in mkpage reldocpath; do
    echo "Building $PROGNAME dist/linix-amd64"
  env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    echo "Building $PROGNAME dist/maxosx-amd64"
  env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    echo "Building $PROGNAME dist/raspberrypi-arm6"
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    echo "Building $PROGNAME dist/raspberrypi-arm7"
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    echo "Building $PROGNAME dist/windows-amd64"
  env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
done

zip -r "$RELEASE_NAME-binary-release.zip" README.md INSTALL.md LICENSE dist/*
