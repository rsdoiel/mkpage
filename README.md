
This is hypothetical at this stage...


# mkpage

An experimental template engine with an embedded markdown processor.  *mkpage* (pronounced "make page") is 
a simple command line tool which accepts key value pairs and applies them to a Golang text/template.
The key side of a pair corresponds to the template keys in the template document (e.g. 
{{.pageContent}} is represented by the key *pageContent*). The value side of the pair can be a string, 
filename or URL for a data source. Here's a simple example of a form letter

```template
    Date: {{- .now}}

    Hello {{.name -}},
    
    The current weather is

    {{.weather}}

    Thank you

    {{.signature}}
```

Render the template above (i.e. myformletter.template) would be accomplished from the following
data sources--

+ "now" and "name" are strings
+ "weather" comes from a URL
+ "signature" comes from a file in our local disc

That would be expressed on the command line as follows

```shell
    mkpage "now=string:$(date)" \
        "name=string:Little Frieda" \
        "weather=http://forecast.weather.gov/MapClick.php?lat=9.9667&lon=139.6667&FcstType=json" \
        signature=testdata/signature.txt \
        testdata/myformletter.template
```

Since we are leveraging Go's [text/template](https://golang.org/pkg/text/template/) the template itself
and be more than simple substitution.

## Options

In additional to populating a template with values from data sources *mkpage* also includes the
[blackfriday](https://github.com/russross/blackfriday) markdown processor.  Using the "-m" option any
a filename referenced with the extension of ".md" will run through the markdown process for being put into 
the template.  This allows you to easily generate pages and website from markdown files using simple templates.

+ -m, -markdown - use a markdown processor when reading ".md" files 
+ -h, -help - get command line help
+ -v, -version - show *mkpage* version number
+ -l, -license - show *mkpage* license information


