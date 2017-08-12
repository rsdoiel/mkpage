
# Action items

## Next

+ [ ] After April 17, 2017 use the new NOAA weather API for example in help and README.md
+ [ ] remove the default template and ship distribution with a set of standard templates
+ [ ] stdin should expect a template, this way _mkpage_ command can be its own pre-processor

## Someday, Maybe

+ [ ] Look at Caddy's Fast-CGI implementation and evaluate integrating a similar approach into _ws_
+ [ ] Add a general purpose indexer that can process both Markdown files and metadata in JSON documents with same name (e.g. page.md, page.json)
+ [ ] Add support to _ws_ to integrate Bleve searches natively given an index name(s) and result templates
+ [ ] adds a csvblock that reads in a CSV file and converts to a GFM table like _csv2mdtable_
+ [ ] titleline and byline should use filenames on the command line if provided (not require -i)
+ [ ] Add let's encrypt support for SSL, see article https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html

