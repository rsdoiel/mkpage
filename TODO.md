
# Action items

## Bugs

+ [ ] After April 17, 2018 use the new NOAA weather API for example in help and README.md (or replace with a new example)
+ [x] Copyright year needs updating in source files
+ [x] Need to handle Toml, JSON front matter based on their respective start/end delimiters

## Next (road to v1.0.0)

+ [ ] Read in mkpage.toml, mkpage.json or mkpage.yaml for mkpage config
+ [ ] Evaluate switching from Blackfriday to GoMarkdown
+ [ ] Add support for [Go Markdown](https://github.com/gomarkdown/markdown)
+ [ ] Add support for [MMark](https://github.com/mmarkdown/mmark)
+ [ ] Add support for rendering remarkjs content
+ [ ] Add support for use "markup" in front matter to pick engine
+ [ ] Add support for passing configuration to markup engine from front matter
+ [ ] Figure out how to comingle Markdown, Fountain, remarkjs safely 
+ [ ] mkpage front matter based on library metadata practices, codemeta.json and relavant Scheme.org scheme
    + [ ] `.doi` the DOI associated with a page
    + [ ] `.creator` should be an array of creator info (e.g. ORCID, given_name, family_name)
    + [ ] `.title`
    + [ ] `.date`
    + [ ] `.publishDate`
    + [ ] `.lastmod`
    + [ ] `.description`
    + [ ] `.draft` (bool)
    + [ ] `.keywords`
    + [ ] `.linkTitle`
    + [ ] `.markdup` (e.g. markdown, fountain, blackfriday, maybe remarkjs)
    + [ ] `.series`
    + [ ] `.slug`
    + [ ] `.type` (e.g. post, article, homepage)
    + [ ] `.permalink`  (e.g. resolver URL)
    + [ ] `.language`
    + [x] `.markdown` holds map of settings to pass to the gomarkdown engine
    + [x] `.fountain` holds map of settings to pass to the Fountain 2 engine
    + [x] `.blackfriday`hold map of settings to pass to the old Blackfriday v2 engine
    + [ ] `.remarkjs` holds the settings for our remarkjs engine
+ [ ] mkpage Sitemap support
    + Current sitemap cli is too naive for sites more than a couple dozen pages
    + Need to support possibly nested sitemap XML references
    + Review Hugo's sitemap support for ideas
    + Need some sort of front matter to identify where/if content would show up in sitemap
+ [ ] mkpage slide support needs to align with remarkjs
    + See https://github.com/gnab/remark, https://remarkjs.com and https://github.com/gnab/remark/wiki
    + [x] Change slide splits from `--` to `---` to conform to remarkjs.com behavior
    + [ ] Consider merging mkpage and mkslide (fewer tools less to learn), consider front matter changes
    + [ ] Add support for slide notes delimited by `???`
+ [ ] Add Support for Hugo style Front Matter configuration of BlackFriday and Fountain engines
+ [x] Add cli for extracting front matter
+ [x] *mkpage* should skip over front matter when rendering
+ [x] Add simple redirect support to _ws_
    + [x] via simple CSV file (from target column, to destination column)

## Someday, Maybe

+ [ ] Should **mkpage** support [RMarkdown](https://rmarkdown.rstudio.com/)?
    + See [RMarkdown, Definitive Guide](https://bookdown.org/yihui/rmarkdown/)
    + Requirements
        + Use of exec to pass content to RMarkdown via separate process
        + an R with RMarkdown installed
+ [ ] Should **mkpage** populate a `.Page` variable with page level metadata?
    + See https://gohugo.io/variables/page/ for definitions in Hugo
    + [ ] `.Alaises` aliases to this page (need to clarify this with mkpage's approach)
    + [ ] `.Content` the content itself defined after the front matter
    + [ ] `.Data` data specific to this type of page
    + [ ] `.Date` 
    + [ ] `.Description`
    + [ ] `.Dir`
    + [ ] `.Draft`
    + [ ] `.ExpiryDate`
    + [ ] `.File` see `.File` for sub-fields
    + [ ] `.FuzzyWordCount` (how fuzzy?)
    + [ ] `.MkPage` see `.MkPage` for sub-fields
    + [ ] `.IsHome`
    + [ ] `.IsNode`
    + [ ] `.IsPage` always true (why include it?)
    + [ ] `.IsTranslated` 
    + [ ] `.Keywords`
    + [ ] `.Kind`
    + [ ] `.Language`
    + [ ] `.Lastmod`
    + [ ] `.LinkTitle`
    + [ ] `.Next`
    + [ ] `.NextInSection`
    + [ ] `.OutputFormats`
    + [ ] `.Pages`
    + [ ] `.Permalink`
    + [ ] `.Plain` strip HTML tags (e.g. what you might want to index with Lunr)
    + [ ] `.PlainWords` content stripped for HTML tags and returned as a slice of string ([]string)
    + [ ] `.PreviousInSection`
    + [ ] `.PublishDate`
    + [ ] `.RawContent` (page content without front matter)
    + [ ] `.ReadingTime`
    + [ ] `.Resources`
    + [ ] `.Ref`
    + [ ] `.RelPermalink`
    + [ ] `.Summary`
    + [ ] `.TableOfContents`
    + [ ] `.Title`
    + [ ] `.Translations`
    + [ ] `.TranslationKey`
    + [ ] `.Truncated`
    + [ ] `.Type`
    + [ ] `.UniqueID` an MD5 Checksum of the file path (relative to content root)
    + [ ] `.Weight`
    + [ ] `.WordCount`
+ [ ] Should **mkpage** populate a `.File` variable with file related metadata
    + [ ] `.File.Path` (the path relative to the current working directory, content root)
    + [ ] `.File.LogicalName` (e.g. foo.en.md)
    + [ ] `.File.TranslationBaseName` (e.g. path.Base() in Go minus any language extension (e.g. foo from foo.en.md))
    + [ ] `.File.ContentBaseName` (e.g. same as TranslationalBaseName for compatibility with Hugo, not sure about Leaf bundle support)
    + [ ] `.File.BaseFileName` (e.g. path.Base(), without extention, e.g. foo.en.md -> foo.en)
    + [ ] `.File.Ext` (e.g. path.Ext())
    + [ ] `.File.Lang` (look at how Hugo handles this, probably uses the two letter language code)
    + [ ] `.File.Dir` (e.g. path.Dir() from current working directory, content root)
+ [ ] Should **mkpage** populate a set of `.MkPage` variables accessible to template or another varname populated and with option to override?
    + [ ] `.MkPage.Version` (semver)
    + [ ] `.MkPage.Generator` (generator string like used in RSS 2)
+ [ ] Should **mkpage** support "shortcodes"?
    + Per Hugo videos, "Shortcodes are partials for content"
    + See https://gohugo.io/templates/shortcode-templates/
    + [ ] Hugo's shortcode is built closely on Go templates using `.Get` for args (positional or k/v), `.Inner` for content contained be a start/end shortcode
    + [ ] Hugo's shortcode as delimited by `{{< shortcode_name >}}` or `{{< shortcode_name >}}{{< /shortcode_name >}}`
    + [ ] Hugo will run shortcode through markdown render engine if prefix/suffix is `{{% shortcode %}}`
+ [ ] Review Hugo's style sitemap variables
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

