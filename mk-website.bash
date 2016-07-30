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
    mkpage \
        "navContent=$nav" \
        "pageContent=$content" \
        page.template > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html with shorthand"
mkPage nav.md README.md index.html
echo "Generating install.html with shorthand"
mkPage nav.md INSTALL.md install.html
