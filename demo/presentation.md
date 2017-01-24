
# What is mkslides?

+ A simple Markdown to HTML slide processor
+ Comes with a simple built-in template 
    + just need your Markdown conten
+ You can use your own slide templates
+ A simple command line tool
+ Only one change to standard Markdown to delimit slides

--

# Getting started

+ Type in your content in Markdown
+ Slides are delimited by "--" line
+ A table of contents slide is also generated

--

# This presentation

This presentation is an example of using mkslides

--

# Customization

*mkslides* can be customized by supply a path to 
a CSS file, to a JavaScript file as well as providing a
template.

--

# Customization - Templates

+ The template engine is a simplied one similar to [Hugo](https://gohugo.io)
+ It is based on Go's text template
+ The following fields are provided to the template for each slide
    + .CurNo - the current slide number of the page
    + .PrevNo - the previous slide number (0 if at the beginning)
    + .NextNo - the number of the next slide (0 if at the end)
    + .FirstNo - the first slide number 
    + .LastNo - the number of the last slide
    + .FName - presentation file's basename (e.g. presention for presentation.md)
    + .Title - title of presentation
    + .Heading - the first H tag found in the slide (used when generating a table of contents)
    + .Content - the Markdown content transformed into HTML
    + .CSSPath - if a custom CSS file is preferred, this is the link
    + .JSPath - if custom JavaScript needs to be included, this is the link

--

# Customization - Default Template

Here is an example template

```text
    <!DOCTYPE html>
    <html>
    <head>
        {{if .Title -}}<title>{{- .Title -}}</title>{{- end}}
        {{if .CSSPath -}}
    <link href="{{ .CSSPath }}" rel="stylesheet" />
       {{else -}}
    <style>
        body {
            width: 100%;
            height: 100%;
            margin: 10%;
            padding: 0;
            font-size: calc(2em+1vw);
            font-family: sans-serif;
        }
        
        ul {
            list-style: disc;
            text-indent: 0.25em;
        }
        
        nav {
            position: absolute;
            top: 0em; 
            margin:0;
            padding:0.24em;
            width: 100%;
            height: 4em;
            text-align: left;
            font-size: 60%;
        }
        
        section {
            width: 100%;
            height: auto;
        }
    </style>
    {{- end }}
    </head>
    <body>
        <nav>
    {{ if ne .CurNo .FirstNo -}}
    <a id="start-slide" href="{{printf "%02d-%s.html" .FirstNo .FName}}">Home</a>
    {{- end}}
    {{ if gt .CurNo .FirstNo -}} 
    <a id="prev-slide" href="{{printf "%02d-%s.html" .PrevNo .FName}}">Prev</a>
    {{- end}}
    {{ if lt .CurNo .LastNo -}} 
    <a id="next-slide" href="{{printf "%02d-%s.html" .NextNo .FName}}">Next</a>
    {{- end}}
        </nav>
        <section>{{ .Content }}</section>
    {{with .JSPath}}<script src="{{.}}"></script>{{end}}
    <script>
    (function (document, window) {
        'use strict';
        var start = document.getElementById('start-slide'),
            prev = document.getElementById('prev-slide'),
            next = document.getElementById('next-slide');
        
        
        document.onkeydown = function(e) {
            switch (e.keyCode) {
                /* case 32: */
                case 37:
                // Previous: left arrow
                    if (prev) {
                        prev.click();
                    }
                    break;
                case 39:
                    // Next: right arrow
                    if (next) {
                        next.click();
                    }
                    break;
                case 72:
                case 83:
                    // Home/Start: h, s
                    if (start) {
                        start.click();
                    }
                    break;
            }
        };
    }(document, window));
    </script>
    </body>
    </html>
```

--

# The basics

```
    mkslides presentations.md
```

This generates your presentation HTML using all the defaults

--

# USAGE

```
    mkslides [OPTIONS] [FILENAME]
```

## OPTIONS:


+ -c,-css &mdash; Specify the CSS file to use
+ -h,-help &mdash; display help
+ -j,-js &mdash; Specify a js file to include
+ -l,-license &mdash; display license
+ -template &mdash; Specify an HTML template to use
+ -t,-title &mdash; Presentation title
+ -v,-version &mdash; display version
 
Version v0.0.3

