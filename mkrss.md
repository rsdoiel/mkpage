
# USAGE

    mkrss [OPTION] HTDOCS BLOG_PATH RSS_FILENAME BASE_URL

## SYNOPSIS

mkrss walks the file system to generate a RSS2 file. It assumes that the directory
for BLOG_PATH is is the base directory conforming to 
BLOG_PATH/YYYY/MM/DD/ARTICLE_HTML where YYYY/MM/DD (Year, Month, Day) 
corresponds to the publication date of ARTICLE_HTML.

```
	-c	If non-zero, limit the number of articles in the RSS file
	-e	A colon delimited list of path exclusions
	-h	display help
	-l	display license
	-v	display version
```

## EXAMPLE

```shell
    mkrss htdocs htdocs/myblog htdocs/blog.rss http://blog.example.org
```

This would build an RSS 2 file in htdocs/blog.rss from the articles in
htdocs/myblog/YYYY/MM/DD.

