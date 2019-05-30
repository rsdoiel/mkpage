
# Action items

## Bugs

+ [ ] After April 17, 2018 use the new NOAA weather API for example in help and README.md

## Next

+ [ ] Add support for metadata taken from Namaste (Name as text) in the directory
+ [ ] Add support to byline, title to leverage Namaste
+ [ ] Add support to inline Fountain markup as well as Markdown
    + I'd like mkslides to generate a script from Fountain markup and slides for presentation from Markdown
+ [ ] Evaluate alternatives to Go's default templates
+ [ ] Add configuration adjustment via front matter
+ [ ] Add mapping of front matter to mkpage paraments passed to templates
+ [x] Add simple redirect support to _ws_
    + [x] via simple CSV file (from target column, to destination column)
+ [ ] implement lunr index generation from front mater in markdown pages and the full body text of markdown pages
+ [ ] consider support mmark as well as Blackfriday markdown engines
+ [ ] python wrapper for mkpage and simple demo of a mini CMS implemented with py_dataset and mkpage.
+ [ ] define themes via minimal set templates and a configuration
    + [ ] minimal theme would be two templates index.tmpl and page.tmpl
    + [ ] CSS and other assets could either be inlined or symlink to appropriate locations in the htdoc root of staging site

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

