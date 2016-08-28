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

function RelativePath() {
    # R is our target result path, 
    # e.g. css/site.css with appropriately prefixed "../"
    R="$1"
    # D is the directory of the filepath we're calculating from, 
    # e.g. lesson1/part2/index.html becomes lesson1/part2
    D=$(dirname $2)
    while [ "$D" != "" ] && [ "$D" != "." ]; do
        R="../$R"
        D=$(dirname $D)
    done
    echo "$R"
}
:
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
