#!/bin/bash

START=$(pwd)
cd $(dirname $0)

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
    if [ -f ./bin/mkpage ]; then
        export PATH=bin:$PATH
    fi

    echo "Rendering $html"
    mkpage \
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
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html

echo "Generating CLI docs"
mkpage \
	"title=text:mkpage: An experimental template and markdown processor" \
        "nav=docs/nav.md" \
        "content=docs/index.md" \
	    "sitebuilt=text:Updated $(date)" \
        "copyright=copyright.md" \
        page.tmpl > docs/index.html

for SCRIPT_NAME in $(findfile -s .bash docs); do
    echo "Running how-to/$SCRIPT_NAME" 
    docs/$SCRIPT_NAME
done

echo "Generating How-To documentation"
mkpage \
	"title=text:mkpage: An experimental template and markdown processor" \
        "nav=how-to/nav.md" \
        "content=how-to/index.md" \
	    "sitebuilt=text:Updated $(date)" \
        "copyright=copyright.md" \
        page.tmpl > how-to/index.html

echo "Generating go-template-recipes.html"
MakePage how-to/nav.md how-to/go-template-recipes.md how-to/go-template-recipes.html

for SCRIPT_NAME in $(findfile -s .bash how-to); do
    echo "Running how-to/$SCRIPT_NAME" 
    how-to/$SCRIPT_NAME
done

cd $START
