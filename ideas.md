---
{
    "has_code": false
}
---

# Ideas and questions

+ Add JSON decoding for text strings as well as local files
+ Could one key/value pair on the command line be expanded be the next key/value pair?
    + {{datetime}} key gets replace by {{date}}, {{time}} and those get response by their respective keys and values?
+ I need a recipe book for leveraging Go's text/template package's DSL, most examples assume you're building a go program first, templates second
+ Add ORCID works to BibTeX as recipe for template development
+ Create xlsx2json that will transform an Excel Workbook (sheet or sheets) into a JSON structure suitable to work with *mkpage*.
    + Filter by Sheet Name(s)
    + Filter by Sheet No ranges (e.g. 0-3 sheets)
    + Filter by column ranges (e.g. A:ZA)
    + Filter by row ranges (e.g. 1:1000)
    + JSON export as an array of rows containing objects with column ID (e.g. A, B, C) as attribute name
    + JSON export as an array of columns contianing array row ordered cells (e.g. 0, 1, 2)
    + JSON export as a array of objects where the header row is used as attribute name and cell value as the object value
+ Look at how [caddy](https://github.com/mholt/caddy) and [xerver](https://github.com/alash3al/xerver) to see how FastCGI was implemented, evaluate if that makes sense for _ws_
