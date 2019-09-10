---
{
    "has_code": true
}
---

# The Basics

_mkpage_ uses Go's text/templates for rendering content. This template system was inspired by
simple templates like [Mustache](https://mustache.github.io/) and [Handlebars](http://handlebarsjs.com/).
While Go's templates can be simple the systems lacks documentation.  As a remedy 
I've collected few simple examples here based on my experience developing websites 
with _mkpage_ and _mkslides_.

## Basic element

Like Mustache and Handlebars Go text/templates use double curly brackets to indicate an
element which is to be replace.  A template that says "Hello" would look something like this

```
    {{ define "hello-world.tmpl" }}Hello {{ .World }}{{ end }}
```

We can use the template to say Hello to "Georgina"

```shell
    echo '{{ define "hello-world.tmpl" }}Hello {{ .World }}{{ end }}' > hello-world.tmpl
    mkpage "World=text:Georgina" hello-world.tmpl
```

Running these two command should result in output like

```
    Hello Georgina
```

The line with the `echo` is just creating our template and saving it as the file _hello-world.tmpl_.
In the template the only special part is `{{ .World }}`. ".World" is a variable, as indicated by the initial '.'.
".World" will be replaced by something we define.  In the line with `mkpage` we define the value for ".World". Note
we don't need to prefix "World" with a dot like we did in the template. We use 'text' before the variable name
to indicate the type of object.  The line tells the template to replace `{{ .World }}` with the text "Georgina".
The last part of the command instructs _mkpage_ to use our _hello-world.tmpl_ template.

If we did not include `World=...` with the _mkpage_ command using the _hello-world.tmpl_ template
_mkpage_ would return output like 

```
    Hello <no value>
```

If we included other key/value pairs not mentioned in the template they would be silently ignored. 

If we made a typo in _hello-world.tmpl_ then we would see an error message.


Try the following to get a feel for how key/value pairs work with _mkpage_. The first two will render but display
`Hello <no value>`. The first example fails because no value is provided and the second fails because the value 
provided doesn't match what is in the template (notice the typo "Wrold" vs. "World").  The next one will display 
an error because _text:_ wasn't included on the value side of the key/value pair.  By default _mkpage_ assumes the
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
    echo '{{ define "title-demo.tmpl" }}' > title-demo.tmpl
    echo "{{if .title}}If this title: {{.title}}{{end}}" >> title-demo.tmpl
    echo "{{with .title}}With this title: {{ . }}{{end}}" >> title-demo.tmpl 
    echo '{{ end }}' >> title-demo.tmpl
    mkpage "title=text:This is a title demo" title-demo.tmpl
```

The output should look like

```
    If this title: This is a title demo
    With this title: This is a title demo
```

In the first line with the *if* we use ".title" as the variable, just like ".World" in our first example.
In the second line we can refer to the value as "." because we used the *with* conditional.  
The reason we prefix variable names with dot (period) is because we are actually describing a path 
or context of object relationships. I like to think of the starting dot as "this here" or simply "this".  
So in the "with" line we are saying "with this title do something" up until the `{{end}}`.
We can refer to ".title" simply as "this thing" or `{{ . }}`, which will be replaced with the value
of ".title".

What happens if you run this command?

```shell
    mkpage title-demo.tmpl
```

This produces two empty lines of output. The reason we don't see something like

```
    If this title: <no value>
    With this title: <no value>
```

is because *if* and *with* are conditionally writing the value of title if it has been set.
This becomes a useful tool when you have content that may or may not exist depending on the
page you're processing.


### Template blocks

Go text/templates support defining blocks and rendering them in conjuction with a main template. This is
also supported by *mkpage*. Each template encountered on the command line is added to an array of templates
parsed by the text/template package.  Each template will be executed and the final results will
render to stdout by *mkpage*.

```shell
    mkpage "content=text:Hello World" testdata/page.tmpl testdata/header.tmpl testdata/footer.tmpl
```

Here is what *page.tmpl* would look like

```go
    {{ define "page.tmpl" }}
    {{template "header" . }}

        {{.content}}

    {{template "footer" . }}
    {{ end }}
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

+ text/plain (e.g. "text:" strings and any file expect those having the extension ".md" or ".json")
+ text/markdown (e.g. "markdown:" strings and file extension ".md")
+ application/json (e.g. "json:" strings and file extension ".json")

It also supports three data sources

+ an explicit string (prefixed with a format, e.g. "text:", "markdown:", "json:")
+ a filepath and filename (the default data source)
+ a URL (identified by the URL prefixes http:// and https://)

Content type is evaluated, transformed (if necessary), and sent to the Go text/template.

Create a template called _data-source-demo.tmpl_. It would look like

```
    {{ define "data-source-demo.tmpl" }}
    This is a plain text string: "{{ .string }}"

    Below is a an included file:
    {{ .file }}
    
    Finally below is data from a URL:
    {{ .url }}
    {{ end }}
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

_mkpage_ implements [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown) 
using the [gomarkdown](https://github.com/gomarkdown/markdown) markdown processor.  This is a 
superset of [Markdown](http://daringfireball.net/projects/markdown/) as created by John Gruber.

The markdown processor is invoked for values with the "markdown:" hint prefix, files ending 
in ".md" extension or URL content with the content type returned as "text/markdown" (i.e. 
content with a type of "text/plain" does not use the markdown process and is treated as plain 
text).

