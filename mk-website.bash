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
    	"Title=text:mkpage: Experimental deconstructed content system" \
        "Nav=$nav" \
        "Content=$content" \
        "CSSPath=text:$csspath" \
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
for FNAME in mkpage mkslides sitemapper reldocpath mkrss byline titleline ws urlencode urldecode; do
  echo "Generating $FNAME.html"
  MakePage nav.md $FNAME.md $FNAME.html
done

echo "Generating docs presentation"
cd demo
if [ -f ../bin/mkslides ]; then
    ../bin/mkslides presentation.md
    ../bin/mkslides three-slides.md
else
    mkslides presentation.md
    mkslides three-slides.md
fi
cd ..

echo "Generating theme examples"
for ITEM in demo/one-element demo/simple demo/simple-with-nav; do
    $ITEM/mk-website.bash
done
