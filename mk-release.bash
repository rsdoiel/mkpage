#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM6 and Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
RELEASE_NAME=mkpage
echo "NOTE: this can take a while..."
for PROGNAME in mkpage reldocpath; do
  echo "Building dist/linux-amd64/$PROGNAME"
  env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  echo "Building dist/maxosx-amd64/$PROGNAME"
  env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  echo "Building dist/raspberrypi-arm6/$PROGNAME"
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  echo "Building dist/raspberrypi-arm7/$PROGNAME"
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  echo "Building dist/windows-amd64/$PROGNAME"
  env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
done

# copy etc/*-example to the distribution folder
for ITEM in etc/*-example; do
  if [ -f "$ITEM" ]; then
     if [ ! -d dist/etc ]; then
       mkdir -p dist/etc
     fi
     cp -vR "$ITEM" dist/
  fi
done

# copy the rest of the distribution items
for ITEM in README.md INSTALL.md LICENSE scripts templates; do
  if [ -f "$ITEM" ]; then
    cp -vR "$ITEM" dist/
  fi
done

# zip up the distribution
zip -r "$RELEASE_NAME-release.zip" dist/*
