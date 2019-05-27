
# Action items

## Bugs

+ [ ] After April 17, 2018 use the new NOAA weather API for example in help and README.md

## Next

+ [x] Add simple redirect support to _ws_
    + [x] via simple CSV file (from target column, to destination column)

## Someday, Maybe

+ [ ] Search via Lunrjs
    + [ ] Document ingest
        + [ ] first pass of ingest is to generate/update metadata
            + [ ] read each Markdown file
            + [ ] parse out any front matter as metadata
            + [ ] use document body for full text field
        + [ ] store metadata results in dataset one document per record
        + [ ] generate default frames based on common fields like title, author, pub date
        + [ ] allow additional frames as needed
        + [ ] store parsed metadata in dataset keyed by paths to docs
    +[ ] Indexing
        + [ ] for each frame generte a LunrJS index
    + Search 
        + [ ] An index aware widget for search (i.e. can turn on/off indexes for searchable corpus)
    + [ ] reads markdown document including front matter creating metadata records stored in dataset
    + [ ] search widget support
        + [ ] generate a search widget if JS enabled
        + [ ] Generate a default search page in Markdown for embedded widget
+ [ ] Add support for metadata taken from Namaste (Name as text) in the directory
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

