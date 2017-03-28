
# mkpage templates and Markdown

_mkpage_ uses Go's text/templates for rendering content. This template system was inspired by
simple templates like [Mustache](https://mustache.github.io/) and [Handlebars](http://handlebarsjs.com/).
While Go's templates can be simple the systems lacks documentation.  As a remedy 
I've collected few simple examples here based on my experience developing websites 
with _mkpage_ and _mkslides_.

## Basic element

Like Mustache and Handlebars Go text/templates use double curly brackets to indicate an
element which is to be replace.  If you wanted to replace "Hello World" with "Hello Georgina" then
your Go template would look something like this

```
    Hello {{ .World }}
```

Here's an example of replacing `{{ .World }}` with "Georgina"

```shell
    echo 'Hello {{ .World }}' > hello-world.tmpl
    mkpage "World=text:Georgina" hello-world.tmpl
```

Running these two command should result in output like

```
    Hello Georgina
```

The line with the `echo` is just creating our template and saving it as the file _hello-world.tmpl_.
In the template the only special part is `{{ .World }}`. This indicates the variable "World" will
be replace by something.  In the line with `mkpage` we define the value for ".World". Note
we don't need to prefix "World" with a dot like we did in the template. We just type 
`World=text:Georgina` in quotes.  This tells the template to replace `{{ .World }}` with the text "Georgina".
At the end of the line starting with _mkpage_ well tell it to use _hello-world.tmpl_ for
the template.

If we did not include `World=...` with the _mkpage_ command using the _hello-world.tmpl_ template
_mkpage_ would return output like 

```
    Hello <no value>
```

If we included other key/value pairs not mentioned in the template they would be silently ignored. 

If we made a typo in _hello-world.tmpl_ then we would see an error message.


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

One nice feature of Go's text template is that template elements can be condition. This can
be done using the "if" and "with" template functions. Here's how to show a title conditionally
using the "if" function.

```go
    {{if .title}}And the title is: {{.title}}{{end}}
```

or using "with"

```go
    {{with .title}}{{ . }}{{end}}
```

Let's create a template file with both these statements called _title-demo.tmpl_ and run the 
_mkpage_ command.

```shell
    echo "{{if .title}}If this title: {{.title}}{{end}}" > title-demo.tmpl
    echo "{{with .title}}With this title: {{ . }}{{end}}" >> title-demo.tmpl 
    mkpage "title=text:This is a title demo" title-demo.tmpl
```

The output should look like

```
    If this title: This is a title demo
    With this title: This is a title demo
```

In the first line with the *if* we refer to ".title" as in our first example with ".World".
In the second line we refer to the value as ".".  The reason we prefix variable names with
a dot (period) is because we are actually describing a path or context of object relationships.
I like to think of the starting dot as "this here" or simply "this".  So in the "with" line
We waying "with this title do something" and between that and the part ending in `{{end}}`
we can refer to ".title" simply as "this thing" where `{{ . }}` is replace with the value
of ".title".

What happens if you run this command?

```shell
    mkpage title-demo.tmpl
```

There is two empty lines of output. The reason is we don't see something like

```
    If this title: <no value>
    With this title: <no value>
```

Is because *if* and *with* are conditionally writing the value of title if it has been set.
This becomes a useful tool when you have content that may or may not exist depending on the
page you're processing.


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

*mkpage* understands three content formats

+ text/plain (e.g. "text:" when specifying strings and any file expect those having the extension ".md" or ".json")
+ text/markdown (e.g. "markdown:" when specifying strings, file extension ".md")
+ application/json (e.g. "json:" when specifying strings, file extension ".json")

It also supports three data sources

+ an explicit string (prefixed with a hint, e.g. "text:", "markdown:", "json:")
+ a filepath and filename (the default data source)
+ a URL (identified by the URL prefixes http:// and https://)

Content type is evaluate and if necessary transformed sending it to the Go text/template.

Create a template called _data-source-demo.tmpl_. It would look like

```
    This is a plain text string: "{{ .string }}"

    Below is a an included file:
    {{ .file }}
    
    Finally below is data from a URL:
    {{ .url }}
```

Create a text file named _hello.md_.

```
    # this is a file

    Hello World!
```

Type the following

```shell
    mkpage "string=text:Hello World" "file=hello.md" \
      "url=https://raw.githubusercontent.com/caltechlibrary/mkpage/master/nav.md" \
      data-source-demo.tmpl
```

What do you see?



## A note about Markdown dialect

_mkpage_ implements [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/) 
using the [blackfriday](https://github.com/russross/blackfriday) markdown processor.  This is a 
superset of [Markdown](http://daringfireball.net/projects/markdown/) as created by John Gruber.

The markdown processor is invoked for values with the "markdown:" hint prefix, files ending 
in ".md" extension or URL content with the content type returned as "text/markdown" (i.e. 
content type of "text/plain" mean the markdown process is not run and the content is 
treated as plain 
text).

