
# This is a demo of the codeblock function from tmplfn

This is the import block in Go from mkpage.go

{{codeblock .src 20 34 "golang"}}


Using _mkpage_ might be used as a preprocessor assembling a markdown document in
markdown.


```shell
    mkpage \
        "content=markdown:$(cat codeblock-demo.md | mkpage src=mkpage.go)" \
        page.tmpl > page.html
```

