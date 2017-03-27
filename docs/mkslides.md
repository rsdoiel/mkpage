
# mkslides

## USAGE

    mkslides [OPTIONS] [FILES]

## SYNOPSIS

mkslides converts a Markdown file into a sequence of HTML5 slides.

+ Use Markdown to write your presentation in one file
+ Separate slides by "--" and a new line (e.g. \n versus \r\n)
+ Apply the simple default template or use your own
+ Control Layout and display with HTML5 and CSS

## CONFIGURATION

+ MKSLIDES_CSS - specify the CSS file to include
+ MKSLIDES_JS - specify the JS file to include
+ MKSLIDES_MARKDOWN - the markdown file to process
+ MKSLIDES_PRESENTATION_TITLE - specify the title of the presentation
+ MKSLIDES_TEMPLATES - specify where to find the templates to use 

## OPTIONS

```
	-c	Specify the CSS file to use
	-css	Specify the CSS file to use
	-h	display help
	-help	display help
	-j	Specify the JavaScript file to use
	-js	Specify the JavaScript file to use
	-l	display license
	-license	display license
	-m	Markdown filename
	-markdown	Markdown filename
	-p	Presentation title
	-presentation-title	Presentation title
	-s	display the default template
	-show-template	display the default template
	-t	A colon delimited list of HTML templates to use
	-templates	A colon delimited list of HTML templates to use
	-v	display version
	-version	display version
```

## EXAMPLE

Here's an example of a three slide presentation

```
    Welcome to [mkslides](../)
    by R. S. Doiel, <rsdoiel@caltech.edu>

    --

    # mkslides

    mkslides can generate multiple HTML5 pages from
    one markdown file.  It splits the markdown file
    on each "--" 

    --

    Thank you

    Hope you enjoy [mkslides](https://github.com/caltechlbrary/mkslides)
```


If you save this as presentation.md and run "mkslides presentation.md" it would
generate the following webpages

+ [00-presentation.html](demo/00-presentation.html)
+ [01-presentation.html](demo/01-presentation.html)
+ [02-presentation.html](demo/02-presentation.html)

