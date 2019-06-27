
# Action items

## Bugs

+ [ ] After April 17, 2018 use the new NOAA weather API for example in help and README.md
+ [x] Copyright year needs updating in source files
+ [x] Need to handle Toml, JSON front matter based on their respective start/end delimiters

## Next

+ [ ] Review Hugo's style sitemap variables
+ [ ] Add support for Hugo Front Matter's predefined variables into rendata resolving system (i.e. treat as if they were specified on the command line as key/data pairs)
    + [ ] aliases
    + [ ] audio
    + [ ] date
    + [ ] description
    + [ ] draft (bool)
    + [ ] expiryDate
    + [ ] headless
    + [ ] images
    + [ ] isCJKLanguage
    + [ ] keywords
    + [ ] layout
    + [ ] lastmod
    + [ ] linkTitle
    + [ ] markdup (We'll support Markdown and Fountain add others later)
    + [ ] outputs (output formats)
    + [ ] publishDate
    + [ ] resources
    + [ ] series
    + [ ] slug
    + [ ] summary
    + [ ] title
    + [ ] type
    + [ ] url
    + [ ] videos
    + [ ] weight
+ [ ] Add Support for Hugo style User Defined Front Matter variables
    + Values are passed as ".Params.\*"
+ [ ] Add Support for Hugo style Front Matter configuration of BlackFriday (and Fountain) engines
+ [ ] mkpage populate a set of `.MkPage` variables accessible to template
    + [ ] `.MkPage.Version` (semver only)
    + [ ] `.MkPage.Environment` (target environment, e.g. macOS, Windows)
    + [ ] `.MkPage.CommitHash` (git commit hash of mkpage)
    + [ ] `.MkPage.BuildDate` (build date of mkpage)
    + [ ] `.MkPage.Generator` (generator string, like defined for RSS 2)
+ [x] Add cli for extracting front matter
+ [x] *mkpage* should skip over front matter when rendering
+ [x] Add simple redirect support to _ws_
    + [x] via simple CSV file (from target column, to destination column)

## Someday, Maybe

+ [ ] Become more Hugo compatible and friendly
    + [ ] Add Hugo style Shortcode support
    + [ ] Add Hugo style Taxonomy variables
    + [ ] Add Hugo style menu entry properites
    + [ ] Add Hugo style File Variables
    + [ ] Add Hugo style Page Variables
    + [ ] Skip Hugo site variables because **mkpage** is page centric
          and that should be handled by the wrapping Python script
+ [ ] Add a tool to generate and search lunr indexes
+ [ ] Add a HugoLike template support to tmplfn use by mkpage.
+ [ ] make markdown engine (blackfriday v2) configurable from front matter
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

