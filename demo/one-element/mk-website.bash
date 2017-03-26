#!/bin/bash

START=$(pwd)
cd $(dirname $0)

function softwareCheck() {
    for NAME in $@; do
        APP_NAME=$(which $NAME)
        if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$NAME" ]; then
            echo "Missing $NAME"
            exit 1
        fi
    done
}

echo "Checking necessary software is installed"
softwareCheck mkpage ws
echo "Converting Markdown files to HTML"
for MARKDOWN_FILE in $(find . -type f | grep -E '.md'); do
    # Calculate the HTML filename
    HTML_FILE="$(dirname $MARKDOWN_FILE)/$(basename $MARKDOWN_FILE .md).html"
    mkpage \
    "Content=$MARKDOWN_FILE" \
    page.tmpl > $HTML_FILE

done

cd $START

