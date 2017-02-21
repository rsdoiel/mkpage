
# mkpage

_mkpage_ (pronounced "make page") project is an experiment in decomposing web content management
systems functionality into a series of simple command line tools -- [mkpage](mkpage.html), 
[mkslide](mkslide.html), [mkrss](mkrss.html), [sitemapper](sitemapper.html), 
[byline](byline.html), [titleline](titleline), [reldocpath](reldocpath.html), 
[slugify](slugify.html), and [ws](ws.html). This makes creating static websites simple.  
Complex sites can be created with a little bit of Bash scripting.

This project started with the *mkpage* command.

## mkpage command

The *mkpage* command is an experimental template engine with an embedded markdown processor. *mkpage* is 
a simple command line tool which accepts key value pairs and applies them to a 
Golang [text/template](https://golang.org/pkg/text/template/).  The key side of a pair corresponds to the 
template element names that will be replaced in the render version of the document. If a key was cllaed
"pageContent" the template element would look like `{{ .pageContent }}`. The value of "pageContent" would
replace `{{ .pageContent }}`. Go text/templates elements can do more than that but the is the core idea.
On the value side of the key/value pair you have strings of one of three formats - plain text, markdown
and JSON.  These three formatted strings can be explicit strings, data from a file or content received from
a URL. Here's a basic demonstration starting with the template.

```template
    Date: {{ .now}}

    Hello {{ .name -}},
    
    Forecast:

    {{range .weather.data.text}}
       + {{ . }}
    {{end}}

    Thank you

    {{.signature}}
```

To render the template above (i.e. myformletter.tmpl) is expecting values from various data sources.
This break down is as follows.

+ "now" and "name" are explicit strings
+ "weather" comes from a URL of JSON content
+ "signature" comes from a file in our local disc

Here is how we would express the key/value pairs on the command line.

```shell
    mkpage "now=text:$(date)" \
        "name=text:Little Frieda" \
        "weather=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" \
        signature=testdata/signature.txt \
        testdata/myformletter.tmpl
```

Notice the two explicit strings are prefixed with "text:" (other possible formats are "markdown:", "json:").
Values without a prefix are assumed to be file paths. We see that in testdata/signature.txt.  Likewise the 
weather data is coming from a URL. *mkpage* distinguishes that by the prefixes "http://" and "https://". 
Since a HTTP response contains headers describing the content type (e.g.  "Content-Type: text/markdown") we 
do not require any other prefix. Likewise a filename's extension can give us an inference of the data format 
it contains. ".json" is a JSON document, ".md" is a Markdown document and everything else is just plain text.


Since we are leveraging Go's [text/template](https://golang.org/pkg/text/template/) the template itself
can be more than a simple substitution. It can contain conditional expressions, ranges for data and even
include blocks from other templates.



## Templates

*mkpage* template engine is the Go [text/template](https://golang.org/pkg/text/template/) package. 
Other template systems could be implemented but I'm keeping the experiment simple at this point.

### Conditional elements

One nice feature of Go's text/template DSL is that template elements can be condition. This can
be done using the "if" and "with" template functions. Here's how to show a title conditionally
using the "if" function.

```go
    {{if .title}}And the title is: {{.title}}{{end}}
```

or using "with"

```go
    {{with .title}}{{ . }}{{end}}
```

### Template blocks

Go text/templates support defining blocks and rendering them in conjuction with a main template. This is
also supported by *mkpage*. For each template encountered on the command line it is added to an array of templates
parsed by the text/template package.  Collectively they are then executed which causes final results 
render to stdout by *mkpage*.

```shell
    mkpage "content=text:Hello World" testdata/page.tmpl testdata/header.tmpl testdata/footer.tmpl
```

Here is what *page.tmpl* would look like

```go
    {{template "header" . }}

        {{.content}}

    {{template "footer" . }}
```

The header and footer are then defined in their own template files (though they also could be combined into one
or even be defined in the main template file itself).

*header.tmpl*

```go
    {{define "header"}}This is the document header{{end}}
```

*footer.tmpl*

```go
    {{define "footer"}}This is the footer{{end}}
```

In this example the output would look like

```text
    This is the document header

        Hello World

    This is the footer
```


## Content formats and data sources

*mkpage* support three content formats

+ text/plain (e.g. "text:" when specifying strings and any file expect those having the extension ".md" or ".json")
+ text/markdown (e.g. "markdown:" when specifying strings, file extension ".md")
+ application/json (e.g. "json:" when specifying strings, file extension ".json")

It also supports three data sources

+ an explicit string (prefixed with a hint, e.g. "text:", "markdown:", "json:")
+ a filepath and filename
+ a URL

Content type is evaluate and if necessary transformed before going into the Go text/template.


## A note about Markdown dialect

In additional to populating a template with values from data sources *mkpage* also includes the
[blackfriday](https://github.com/russross/blackfriday) markdown processor.  The `blackfriday.MarkdownCommon()`
function is envoked whenever markdown content is suggested. That means for strings that have the 
"markdown:" hint prefix, files ending in ".md" file extension or URL content with the content type
returned as "text/markdown".


## Options

+ -h, -help - get command line help
+ -v, -version - show *mkpage* version number
+ -l, -license - show *mkpage* license information
+ -t, -template - show *mkpage*'s default template


## Companion utilities

*mkpage* comes with some helper utilities that make scripting a deconstructed
content management system from Bash easier.

### mkslides

*mkslides* generates a set of HTML5 slides from a single Markdown file. It uses
the same template engine as *mkpage*

### mkrss

*mkrss* will scan a directory tree for Markdown files and add each markdown file with
a corresponding HTML file to the RSS feed generated.

### byline

*byline* will look inside a markdown file and return the first _byline_ it finds
or an empty string if it finds none. The _byline_ is identified with a regular
expression. This regular expression can be overriden with a command line option.

### titleline

*titleline* will look inside a markdown file and return the first h1 exquivalent title
it finds or an empty string if it finds none. 

### reldocpath

*reldocpath* is intended to simplify the calculation of relative
asset paths (e.g. common css files, images, feeds) when working from
a common project directory.

#### Example

You know the path from the source document to target document from the project root folder.

+ Source is *course/week/01/readings.html*  
+ Target is *css/site.css*.

In Bash this would look like--

```shell
    # We know the paths relative to the project directory
    DOC_PATH="course/week/01/readings.html"
    CSS_PATH="css/site.css"
    echo $(reldocpath $DOC_PATH $CSS_PATH)
```

the output would look like

```shell
    ../../../css/site.css
```

### slugify

*slugify* takes one or more command line args (e.g. a phrase like "Hello World") and return
an updated version that is more friendly for filenames and URLS (e.g. "Hello-World").

#### Example

```shell
    slugify My thoughts on functional programming
```

Would yield

```
    My-thoughts-on-functional-programming
```

### ws

*ws* is a simple static file webserver.  It is suitable for viewing your local copy
of your static website on your machine.  It runs with minimal resources and by default
will serve content out to the URL http://localhost:8000.  It can also be used to host
a static website and has run well on small Amazon virtual machines as well as Raspberry Pi
computers acting as local private network web servers.

#### Example

```shell
    ws Sites/mysite.example.org
```

This would start the webserver up listen for browser requests on _http://localhost:8000_.
The content viewable by your web browser would be the files inside the _Sites/mysite.example.org_
directory.

```shell
    ws -url http://mysite.example.org:80 Sites/mysite.example.org
```

Assume the machine where you are running *ws* has the name mysite.example.org then your could
point your web browser at _http://mysite.example.org_ and see the web content you have in 
_Site/mysite.example.org_ directory.
