#!/bin/bash

function checkApp() {
    APP_NAME=$(which $1)
    if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in $@; do
        checkApp $APP_NAME
    done
}

function MakePage () {
    nav="$1"
    content="$2"
    html="$3"
    # Always use the latest compiled mkpage
    APP=$(which mkpage)
    if [ -f ./bin/mkpage ]; then
        APP="./bin/mkpage"
    fi

    echo "Rendering $html"
    $APP \
	"title=text:mkpage: An experimental template and markdown processor" \
        "nav=$nav" \
        "content=$content" \
	    "sitebuilt=text:Updated $(date)" \
        "copyright=copyright.md" \
        page.tmpl > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating go-template-recipes.html"
MakePage nav.md go-template-recipes.md go-template-recipes.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html
echo "Generating mkslides.html"
MakePage nav.md mkslides.md mkslides.html
echo "Generating reldocpath.html"
MakePage nav.md reldocpath.md reldocpath.html
echo "Generating slugify.html"
MakePage nav.md slugify.md slugify.html
echo "Generating mkpage.html"
MakePage nav.md mkpage.md mkpage.html
echo "Generating sitemapper.html"
MakePage nav.md sitemapper.md sitemapper.html
