
# USAGE

	mkslides [OPTIONS] [KEY/VALUE DATA PAIRS] MARKDOWN_FILE [TEMPLATE_FILENAMES]

## DESCRIPTION



SYNOPSIS

mkslides converts a Markdown file into a sequence of HTML5 slides using the
key/value pairs to populate the templates and render to stdout.

Features

+ Use Markdown to write your presentation in one file
+ Separate slides by "--" and a new line (e.g. \n versus \r\n)
+ Apply the default template or use your own
+ Control Layout and display with HTML5 and CSS

mkslides is based on mkpage with the difference that multiple pages
result from a single Markdown file. To manage the linkage between
slides some predefined template variables is used.

+ "title" which would hold the page title for presentation
+ "header" which would hold a header section for the presentation (e.g. organization logo)
+ "footer" which would hold a footer section for the presentation (e.g. copyright statement)
+ "nav" which would hold an alternative navigation section for the presentation
+ "csspath" which would hold the path to your CSS File.
+ "content" holds the extracted for each slide
+ "cur_no" which holds the current page number
+ "first_no" which holds the first slide's page number (e.g. 00)
+ "last_no" which holds the last slides page number (e..g length of slide deck minus one)
+ "prev_no" which holds the previous slide number if CurNo is create than 0
+ "next_no" which holds the next slide number if CurNo is not the last slide
+ "filename" is the filename for presentation

In your custom templates these should be exist to link everything together
as expected.  In addition you may want to include JavaScript to allow mapping
actions like "next slide" to the space bar or mourse click.


## ENVIRONMENT

Environment variables can be overridden by corresponding options

```
    MKPAGE_TEMPLATES   # a colon delimiter list of default templates to use
```

## OPTIONS

Below are a set of options available. Options will override any corresponding environment settings.

```
    -c                   Specify the CSS file to use
    -css                 Specify the CSS file to use
    -example             display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h                   display help
    -help                display help
    -j                   Specify the JavaScript file to use
    -js                  Specify the JavaScript file to use
    -l                   display license
    -license             display license
    -m                   Markdown filename
    -markdown            Markdown filename
    -p                   Presentation title
    -presentation-title  Presentation title
    -s                   display the default template
    -show-template       display the default template
    -t                   A colon delimited list of HTML templates to use
    -templates           A colon delimited list of HTML templates to use
    -v                   display version
    -version             display version
```


## EXAMPLES



EXAMPLE

In this example we're using the default slide template.
Here's an example of a three slide presentation

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

If you saved this as presentation.md you can run the following
command to generate slides

    mkslides "title=text:My Presentation" \
	    "csspath=text:css/slides.css" presentation.md

Using a custom template would look like

    mkslides -t custom-slides.tmpl \
        "title=text:My Presentation" \
	    "csspath=text:css/slides.css" presentation.md

This would result in the following webpages

+ 00-presentation.html
+ 01-presentation.html
+ 02-presentation.html



mkslides v0.0.26
