#!/bin/bash

START="$(pwd)"
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
	for FNAME in $(find . -type f | grep -E '.md'); do
		title="$(titleline -i "$FNAME")"
		docpath="$(dirname "$FNAME")"
		if [ "$docpath" = "./" ]; then
			docpath=""
		fi
		html_filename="$(basename "$FNAME" .md).html"
		if [ "$html_filename" != "nav.html" ] && [ "$html_filename" != "index.html" ]; then
			if [ "$docpath" != "" ]; then
				echo "+ [$title]($docpath/$html_filename)" >> nav.md
			else
				echo "+ [$title]($html_filename)" >> nav.md
			fi
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
	if [ "$FNAME" != "nav.md" ] && [ "$FNAME" != "index.md" ]; then
		# Caltechlate DOCPath
		DOCPath="$(dirname "$MARKDOWN_FILE")"
		# Calculate the HTML filename
		HTML_FILE="$(basename "$MARKDOWN_FILE" .md).html"
		CSSPath="$(reldocpath "$DOCPath" css)"
		WEBSITE_TITLE="$(titleline -i "$MARKDOWN_FILE")"
		echo "Generating $WEBSITE_TITLE, $HTML_FILE from $MARKDOWN_FILE"
		mkpage \
			"title=text:$WEBSITE_TITLE" \
			"CSSPath=text:$CSSPath/site.css" \
			"nav=nav.md" \
			"content=$DOCPath/$FNAME" \
			../page.tmpl > "$DOCPath/$HTML_FILE"
	fi
done

cd "$START"
