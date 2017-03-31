#!/bin/bash

START=$(pwd)
cd "$(dirname "$0")"

function softwareCheck() {
    for NAME in "$@"; do
        APP_NAME=$(which "$NAME")
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
    csspath=$("$RELDOCPATH" "$html" css/site.css)

    echo "Rendering $html"
    "$MKPAGE" \
    	"Title=text:mkpage: Experimental deconstructed content system" \
        "Nav=$nav" \
        "Content=$content" \
        "CSSPath=text:$csspath" \
        page.tmpl > "$html"
}

echo "Checking necessary software is installed"
softwareCheck mkpage mkslides reldocpath
echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html

for FNAME in docs/index docs/mkpage docs/mkslides docs/sitemapper docs/reldocpath docs/mkrss docs/byline docs/titleline docs/ws docs/urlencode docs/urldecode docs/go-template-recipes; do
    D=$(dirname "$FNAME")
    echo "Generating $FNAME.html"
    MakePage "$D/nav.md" $FNAME.md $FNAME.html
done

echo "Generating slides demo"
cd docs/slides
if [ -f ../bin/mkslides ]; then
    ../../bin/mkslides presentation.md
    ../../bin/mkslides three-slides.md
else
    mkslides presentation.md
    mkslides three-slides.md
fi
cd "$START"

echo "Generating theme demos"
for ITEM in docs/one-element docs/simple docs/simple-with-nav; do
    $ITEM/mk-website.bash
done
