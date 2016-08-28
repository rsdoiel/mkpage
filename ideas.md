
# Ideas and questions

+ Add JSON decoding for text strings as well as local files
+ Recipe book for Go temlates
    + Basics of create a template
    + Idomatic template solutions
    + Should some open source/cc0 Jekyll and Hugo themes be ported to CSS and basic Go templates?
+ Additional utilities
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
    + Should these utilities have simple names or share a prefix?
