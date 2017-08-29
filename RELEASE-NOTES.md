
# Release Notes

## v0.0.18

+ Templates are now all assumed to start with a define with the master template listed first and matching its basename
    + Affects _mkpage_, _mkslides_
+ Various bug fixes
    + Fixed some CORS handling in _ws_
+ Added ACME cert support for https in _ws_
    + Any http request will automatically redirect to https when ACME cert support is enabled
+ mkpage, mkslide accept templates via stdin, normalize key/value data pairs between them, updated mkslide docs
+ Bug fixes
    + Improved documentation
    + Fixed sitemapper bug when mapping from the current working directory
    + Update copyright link for Caltech Library

