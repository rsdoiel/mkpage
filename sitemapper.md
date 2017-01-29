
# USAGE

    sitemapper [OPTIONS] HTDOCS_PATH SITEMAP_FILENAME PUBLIC_BASE_URL

## OVERVIEW

sitemapper generates a sitemap for the website.

```
	-e	A colon delimited list of path parts to exclude from sitemap
	-exclude	A colon delimited list of path parts to exclude from sitemap
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-u	Set the change frequencely value, e.g. daily, weekly, monthly
	-update-frequency	Set the change frequencely value, e.g. daily, weekly, monthly
	-v	display version
	-version	display version
```

## EXAMPLE

```shell
    sitemapper htdocs htdocs/sitemap.xml http://blog.example.org
```

This yields a sitemap.xml file in the htdocs folder.
