#!/bin/bash

function checkApp() {
    APP_NAME=$1
    if [ "$APP_NAME" = "" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in $@; do
        checkApp $APP_NAME
    done
}

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html from $content and $nav"
    mkpage -m \
	"title=string:mkpage: A hypothetical template and markdown processor" \
        "nav=$nav" \
        "content=$content" \
	    "sitebuilt=string:Updated $(date)" \
        "copyright=copyright.md" \
        page.template > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html with mkpage"
mkPage nav.md README.md index.html
echo "Generating install.html with mkpage"
mkPage nav.md INSTALL.md install.html
