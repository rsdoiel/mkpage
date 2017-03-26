#!/bin/bash

function SoftwareCheck() {
    for NAME in $@; do
        APP_NAME=$(which $NAME)
        if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$NAME" ]; then
            echo "Missing $NAME"
            exit 1
        fi
    done
}

echo "Checking necessary software is installed"
SoftwareCheck mkpage reldocpath ws
if [ "$WEBSITE_TITLE" = "" ]; then
    WEBSITE_TITLE="Simple Theme Demo"
fi

echo "Converting Markdown files to HTML supporting a relative document path to the CSS file"
for MARKDOWN_FILE in $(find . -type f | grep -E '.md'); do
    # Caltechlate DOCPath
    DOCPath=$(dirname $MARKDOWN_FILE)
    # Calculate the HTML filename
    HTML_FILE="$DOCPath/$(basename $MARKDOWN_FILE .md).html"
    CSSPath=$(reldocpath $DOCPath css)
    mkpage \
        "Title=text:$WEBSITE_TITLE" \
        "CSSPath=text:$CSSPath/site.css" \
        "Content=$MARKDOWN_FILE" \
        page.tmpl > $HTML_FILE
done
