
# mkpage

## USAGE

    mkpage [OPTION] [KEY/VALUE DATA PAIRS] TEMPLATE_FILENAME [TEMPLATE_FILENAMES]

## SYNOPSIS

Using the key/value pairs populate the template(s) and render to stdout.

## CONFIGURATION

You can set a local default template path by using environment variables.

+ MKPAGE_TEMPLATES - is the colon delimited list of template paths

## OPTIONS

```
	-h	show help
	-help	show help
	-l	show license
	-license	show license
	-s	display the default template
	-show-template	display the default template
	-t	colon delimited list of templates to use
	-templates	colon delimited list of templates to use
	-v	show version
	-version	show version
```

## EXAMPLE

Template

```
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
 + "weatherForcast" comes from a URL
 + "license" comes from a file in our local disc

That would be expressed on the command line as follows

```shell
	mkpage "now=text:$(date)" "name=text:Little Frieda" \
		"weather=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" \
		signature=testdata/signature.txt \
		testdata/myformletter.template
```

Golang's text/template docs can be found at 

```
      https://golang.org/pkg/text/template/
```

