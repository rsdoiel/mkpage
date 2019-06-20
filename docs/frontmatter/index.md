
# USAGE

	frontmatter [OPTIONS]

## DESCRIPTION


frontmatter extracts a front matter from a Markdown file.
If no front matter is present then an empty file 
is returned. NOTE: frontmatter doesn't process the front 
matter it only extracts it.


## OPTIONS

Below are a set of options available.

```
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           input filename
    -l, -license         display license
    -o, -output          output filename
    -quiet               suppress error messages
    -v, -version         display version
```


## EXAMPLES


Extract a front matter from article.md.

    cat article.md | frontmatter

This will display the front matter if found in article.md.

    frontmatter -i article.md

Will also do the same.


frontmatter v0.0.25-rsdoiel
