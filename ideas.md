
# Ideas and questions

+ Recipe book for Go temlates
    + Basics of create a template
    + Idomatic template solutions
    + Leveraging API sources
            + ORCID Profile to CV Web Page use case
            + ORCID Profile to Pubs List in BibTeX and Web Page use case
            + ORCID Pubs list as RSS Feed
    + Should some open source/cc0 Jekyll and Hugo themes be ported to CSS and basic Go templates?
+ Additional utilities
    + idxpage
        + given a markdown file and Bleve index file name add markdown/html file to index
    + title
        + given a markdown file, return the first "#" value
        + Have an option to support picking out the nth "#"
    + byline
        + given a markdown file, pluckout a byline based in a prefix or regexp string
    + relpath
        + given two paths with the same base directory
        + generate a relative path from one to another
    + breadcrumb
        + given a path
        + render as JSON 
            + unslug the path elements
            + associate with the approprate relative URL
        + organize the JSON so it is easy to process in a template
    + menu from subdirectories
        + make a list of subdirections
        + unslug the names
        + associate with appropriate relative URL
        + How to you autogenerate but support re-ordering without hardcoding directory sort orders
            + do you supply some sort order function?
+ Should their be any integration of *mkpage* with a Bleve Index and search service?
    + How much metadata can I get from path and document contents (without resorting to Front Matter)?
    + Would this be enough to populate a BibFrame record?

## Someday, Maybe ideas

+ autotag
    + automatically collect tag words from Markdown documents
        + tags can be identified by quotes, *emphasis*, _underscores_ and link labels
    + should render a JSON structure for building tag clouds and tag indexes
+ wssearch/wsindex
    + Indexer could be from JSON blobs that conform to a subset of schema.org definitions
    + A generic search engine could implemented by reading indexes and accepting a map of search fields/constraints (mirroed in the supported schema)

    
