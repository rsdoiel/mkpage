
# Action items

## Bugs

+ [ ] After April 17, 2018 use the new NOAA weather API for example in help and README.md

## Next

+ [ ] Add support for metadata taken from Namaste (Name as text) in the directory

## Someday, Maybe

+ [ ] Add redirect support to _ws_ via some sort of JSON object
+ [ ] Align templating and feature set with Hugo while retaining the approach of simple commands responsible for simple actions in a pipe line
+ [ ] Remove the default template and ship distribution with a set of standard templates
+ [ ] Add support for integration with pygments for syntax highlighting
+ [ ] Align tmplfn with Hugo's extensions on to Go template (e.g. isset)
+ [ ] Add front matter support
+ Investigate moving beyond Go templates 
    + [ ] review [Mustache](https://github.com/hoisie/mustache) or Handlebars via [velvet](https://github.com/gobuffalo/velvet) or [raymond](https://github.com/aymerick/raymond)
    + [ ] review [ace](https://github.com/yosssi/ace) templates
    + [ ] review [amber](https://github.com/eknkc/amber) templates
    + [ ] review [Pango2](https://github.com/flosch/pongo2) templates
    + [ ] review [Liquid](https://github.com/osteele/liquid) templates
    + [ ] Add support for [Damsel](https://github.com/dskinner/damsel)
+ [ ] Configurable Mime-Type assignments
+ [ ] Look at Caddy's Fast-CGI implementation and evaluate integrating a similar approach into _ws_
+ [ ] Add a general purpose indexer that can process both Markdown files and metadata in JSON documents with same name (e.g. page.md, page.json)
+ [ ] Add support to _ws_ to integrate Bleve searches natively given an index name(s) and result templates
+ [ ] Adds a csvblock that reads in a CSV file and converts to a GFM table like _csv2mdtable_
+ [ ] titleline and byline should use filenames on the command line if provided (not require -i)

