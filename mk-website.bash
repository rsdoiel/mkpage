#!/bin/bash

function softwareCheck() {
    for NAME in $@; do
        APP_NAME=$(which $NAME)
        if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$NAME" ]; then
            echo "Missing $NAME"
            exit 1
        fi
    done
}

function MakePage () {
    nav="$1"
    content="$2"
    html="$3"
    # Always use the latest compiled mkpage and reldocpath
    MKPAGE=$(which mkpage)
    if [ -f ./bin/mkpage ]; then
        MKPAGE="./bin/mkpage"
    fi
    RELDOCPATH=$(which reldocpath)
    if [ -f ./bin/reldocpath ]; then
        RELDOCPATH="./bin/reldocpath"
    fi
    csspath=$($RELDOCPATH $html css/site.css)

    echo "Rendering $html"
    $MKPAGE \
	"title=text:mkpage: An experimental template and markdown processor" \
        "nav=$nav" \
        "content=$content" \
        "csspath=text:$csspath" \
        page.tmpl > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage reldocpath
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

echo "Generating docs presentation"
if [ -f bin/mkslides ]; then
    ./bin/mkslides presentation.md
else
    mkslides presentation.md
fi


