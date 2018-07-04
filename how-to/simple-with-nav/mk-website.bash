#!/bin/bash

START=$(pwd)
cd "$(dirname "$0")"

function SoftwareCheck() {
	for NAME in "$@"; do
		APP_NAME="$(which "$NAME")"
		if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$NAME" ]; then
			echo "Missing $NAME"
			exit 1
		fi
	done
}

function GenerateNav() {
	echo "+ [Home](/)" >nav.md
	echo "+ [Up](../)" >>nav.md
	for FNAME in $(find . -type f | grep -E "\.md$"); do
		title="$(titleline -i "$FNAME")"
		docpath="$(dirname "$FNAME")"
		html_filename="$(basename "$FNAME" .md).html"
		if [ "$html_filename" != "nav.html" ]; then
			echo "+ [$title]($docpath/$html_filename)" >>nav.md
		fi
	done
}

echo "Checking necessary software is installed"
SoftwareCheck mkpage reldocpath titleline ws

echo "Generating nav.md from current Markdown files found"
GenerateNav

echo "Converting Markdown files to HTML supporting a relative document path to the CSS file"
for MARKDOWN_FILE in $(find . -type f | grep -E "\.md$"); do
	# Get filename
	FNAME="$(basename "$MARKDOWN_FILE")"
	if [ "$FNAME" != "nav.md" ]; then
		# Caltechlate DOCPath
		DOCPath="$(dirname "$MARKDOWN_FILE")"
		# Calculate the HTML filename
		HTML_FILE="$(basename "$MARKDOWN_FILE" .md).html"
		CSSPath="$(reldocpath "$DOCPath" css)"
		WEBSITE_TITLE="$(titleline -i "$MARKDOWN_FILE")"
		echo "Generating $WEBSITE_TITLE from $MARKDOWN_FILE"
		mkpage \
			"Title=text:$WEBSITE_TITLE" \
			"CSSPath=text:$CSSPath/site.css" \
			"Nav=nav.md" \
			"Content=$MARKDOWN_FILE" \
			page.tmpl >"$DOCPath/$HTML_FILE"
	fi
done

cd "$START"
