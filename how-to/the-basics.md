
# mkpage templates and Markdown

_mkpage_ uses Go's text/templates for rendering content. This template system was inspired by
simple templates like [Mustache](https://mustache.github.io/) and [Handlebars](http://handlebarsjs.com/).  
While Go's take on these template systems is simple it lacks documentation.  To start to remedy 
that I've collected the basics used I've commonly in the websites I've built using _mkpage_ and _mkslides_.

## Basic element

Like Mustache and Handlebars Go text/templates using a double curly bracket notation to indicate an
element to replace.  If you wanted to place "Hello World" with "Hello Georgina" then
your Go template would look something like this

```
    Hello {{ .World }}
```

Here's an example of replacing {{ .World }} would get replaced with "Georgina" using _mkpage_

```shell
    echo 'Hello {{ .World }}' > hello-world.tmpl
    mkpage "World=text:Georgina" hello-world.tmpl
```

To he output should look like

```
    Hello Georgina
```

The line with the `echo` is just creating our template and saving it as the file _hello-world.tmpl_.
In the template the only special part is that `{{ .World }}` indicating the variable "World" will
be replace by something.  In the line with `mkpage` we are define the value for ".World". Note
we don't need to prefix "World" with a dot. We just type `World=text:Georgina` in quotes.
This tells the template to replace `{{ .World }}` with the text "Georgina".
The end of the line with _mkpage_ simply says use the template _hello-world.tmpl_ we previously created.

If we did not include `World=...` wth the _mkpage_ command using the _hello-world.tmpl_ template
_mkpage_ would return an error. If we included other key/value pairs not mentioned in the template
they would be silently ignored. 

Try the following to get a feel for how key/value pairs work with _mkpage_. The first two will render but display
`Hello <no value>` where "Georgina" was in our previous example. The first one because no value is
provided and the second one because the value provided doesn't match what is in the template (i.e.
notice the typo "Wrold" vs. "World").  The next one will display an error because
_text:_ wasn't included on the value side of the key/value pair.  By default _mkpage_ assumes the
value is refering to a file and in this case can't find the file Georgina in your current directory.  
The last two will display `Hello Georgina` should display since the value for "World" is provided. The
last one just ignores "Name=text:Fred" because name isn't found in the template.

```shell
    mkpage hello-world.tmpl
    mkpage "Wrold=text:Georgina" hello-world.tmpl
    mkpage "World=Georgina" hello-world.tmpl
    mkpage "World=text:Georgina" "Name=text:Fred" hello-world.tmpl
    mkpage "World=text:Georgina" hello-world.tmpl
```


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

