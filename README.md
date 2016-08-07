
This is experimental..., things are sure to change

# mkpage

An experimental template engine with an embedded markdown processor.  *mkpage* (pronounced "make page") is 
a simple command line tool which accepts key value pairs and applies them to a Golang [text/template](https://golang.org/pkg/text/template/).
The key side of a pair corresponds to the template keys in the template document (e.g. 
{{.pageContent}} is represented by the key *pageContent*). The value side of the pair can be a string, 
filename or URL for a data source. Here's a simple example of a form letter

```template
    Date: {{.now}}

    Hello {{.name -}},
    
    Forecast:

    {{range .weather.data.text}}
       + {{ . }}
    {{end}}

    Thank you

    {{.signature}}
```

Render the template above (i.e. myformletter.template) would be accomplished from the following
data sources--

+ "now" and "name" are strings
+ "weather" comes from a URL of JSON content
+ "signature" comes from a file in our local disc

That would be expressed on the command line as follows

```shell
    mkpage "now=text:$(date)" \
        "name=text:Little Frieda" \
        "weather=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" \
        signature=testdata/signature.txt \
        testdata/myformletter.template
```

Since we are leveraging Go's [text/template](https://golang.org/pkg/text/template/) the template itself
can be more than a simple substitution.

## Template blocks

The Go text templates support defining blocks and rendering them in conjuction with a main template. This is
also supported by *mkpage*. For each template encountered on the command line it is added to an array of templates
passed and parse by the text template package.  This is then executed and output rendered by *mkpage*.

```shell
    mkpage "content=text:Hello World" testdata/page.tmpl testdata/header.tmpl testdata/footer.tmpl
```

Here is what *page.tmpl* would look like

```go
    {{template "header" . }}

        {{.content}}

    {{template "footer" . }}
```

The header and footer are then defined in their own template files (though they also could be combined into one).

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

+ text/plain (e.g. "text:" when specifying strings, any file extension except ".md", and ".json")
+ text/markdown (e.g. "markdown:" when specifying strings, file extension ".md")
+ application/json (e.g. "json:" when specifying strings, file extension ".json")

It also supports three content sources

+ an explicit string (prefixed with a type, e.g. "text:", "markdown:", "json:")
+ a filepath
+ a URL


## Templates

*mkpage* template engine is the Go [text/template](https://golang.org/pkg/text/template/) package. 
Other template systems could be implemented but I'm keeping the experiment simple at this point.


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


