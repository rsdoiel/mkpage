
# Go text/template recipes

*mkpage* template engine is Go's [text/template](https://golang.org/pkg/text/template/). Go's templates
provide a flexible and simple [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) describing
how to assemble a document based on a data structure passed to it.  *mkpage* uses a list of key/value
pairs on the command line to populate the data structure the template package expects.  This includes
support for JSON formatted text from strings, files and URL response.

While Go's template package is not complicated to use it doesn't come with allot of example or tutorials.
Most articles you find on Go's template package either are for writing your own server code using the
templates to format responses or their for sofisticated content generation systems like [Hugo](http://gohugo.io)
that extend the basic template system by large measures towards systems like Jykell or Ruby's Jade templates.

*mkpage* uses Go's template as is.  It is meant to be a trivially easy system for producing simple content whether 
plain text or HTML.

## simple data map with get well card

In this example we want to add a name to a simple get well message.

Our template is called **get-well.tmpl**. It looks like 

```go
    Dear {{ .name -}},

    Hope you are feeling better today.

    Sencerly,

    Mojo Sam
```

On the command line we can run *mkpage* with the following options

```shell
    mkpage "name=string:Little Frieda" get-well.tmpl
```

The output would look like

```text
    Little Frieda,

    Hope you are feeling better today.

    Sencerly,

    Mojo Sam
```

### Explanation

The key "name" has a string value of "Little Frieda".  The template indicates this needs to be included 
after the word "Dear". The key "name" is proceeded by a period or dot.  The substitution happens between 
the opening "{{" and closing "}}".  Notice the "-" before the closing "}}". This tells the template 
engine to not allow spacas after the value and the next non-space character (i.e. the comma of the 
opening line).

## JSON data a key/value blob report

In this example we construct a JSON object as part of the key/value pairs on the command line and
pass it through the blob.tmpl template that displays they pairs.

The command envokation looks like

```shell
    mkpage 'blob=json:{"one":1,"two":2}'  blob.tmpl
```

The template is a simple range construct

```go
    {{range $key,$val := .blob }}
        Key: {{ $key }} Value: {{ $val -}}
    {{end}}
```

Results in text like

```text
    
       Key: one Value: 1
       Key: two Value: 2

```

### Explanation

We use the range function to iterate over the key/value pairs of our JSON object. Additionally
we assign those values to the template variables called "$key" and "$val". These are then used
to format our output. Also notice the trailing values "-" which supresses and extra new line.

## JSON data with weather forecast

In this example we get the current weather forecast for Guam.  The source of the weather information
is [NOAA](http://noaa.gov)'s [National Weather Services](http://weather.gov) website.  By including the
parameter "FcstType=json" at the end of the URL you get a JSON version of the weather forecast rather 
than the HTML or XML alternatives.

+ data source: http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FsctType=json

Our template will be call **forecast.tmpl**. It will be used to produce a Markdown file of weather related
information obtained from the JSON response.

```go
    {{with $co := .forecast.currentobservation}}
    Current Observation:

        + {{ $co.name }}
        + Elevation: {{ $co.elev }}
        + Latitude: {{ $co.latitude }}
        + Longitude: {{ $co.longitude }}
        + Date: {{ $co.Date }}
        + Temp: {{ $co.Temp }}
        + Dew Point: {{ $co.Dewp }}
        + Relative Humidity: {{ $co.Relh }}
        + Winds: {{ $co.Winds }}
        + Wind direction: {{ $co.Windd }}
        + Gust: {{ $co.Gust }}
        + Visibility: {{ $co.Visibility }}

    {{end}}

    Forecast:
    {{range .forecast.data.text }}
        + {{ . -}}
    {{end}}
```

The command line for *mkpage* would look like

```shell
    mkpage "forecast=http://forecast.weather.gov/MapClick.php?lat=13.47190933300044&lon=144.74977715100056&FcstType=json" testdata/forecast.tmpl
```

The resulting page should look something like 

```text

    Current Observation:

        + Agana, Guam International Airport
        + Elevation: 299
        + Latitude: 13.48
        + Longitude: 144.8
        + Date: 5 Aug 08:54 am ChST
        + Temp: 82
        + Dew Point: 79
        + Relative Humidity: 89
        + Winds: 12
        + Wind direction: 220
        + Gust: 0
        + Visibility: 10.00


    Forecast:

        + Scattered showers and thunderstorms.  Mostly cloudy, with a high near 84. Breezy, with a southwest wind 23 to 25 mph, with gusts as high as 32 mph.  Chance of precipitation is 40%.
        + Scattered showers and thunderstorms.  Mostly cloudy, with a low around 79. Breezy, with a southwest wind 15 to 20 mph, with gusts as high as 25 mph.  Chance of precipitation is 40%.
        + Mostly cloudy, with a high near 88. Heat index values as high as 99. Breezy, with a southwest wind 17 to 21 mph, with gusts as high as 26 mph. 
        + Mostly cloudy, with a low around 79. Southwest wind 13 to 17 mph, with gusts as high as 22 mph. 
        + Mostly cloudy, with a high near 88. Southwest wind 14 to 17 mph, with gusts as high as 22 mph. 
        + Mostly cloudy, with a low around 79.
        + Mostly sunny, with a high near 89.
        + Partly cloudy, with a low around 80.
        + Scattered showers and thunderstorms.  Mostly cloudy, with a high near 89. Chance of precipitation is 40%.
        + Scattered showers and thunderstorms.  Mostly cloudy, with a low around 79. Chance of precipitation is 40%.
        + Scattered showers and thunderstorms.  Mostly cloudy, with a high near 89. Chance of precipitation is 40%.
        + Scattered showers and thunderstorms.  Mostly cloudy, with a low around 79. Chance of precipitation is 40%.
        + Partly sunny, with a high near 89.
```


