USAGE: slugify [OPTIONS] STRING_TO_SLUGIFY

SYNOPSIS

%!c(string=slugify)hanges a human readable string into a path or URL friendly
string. E.g. "Hello World" becomes "hello-world"

	-h	display help
	-l	display license
	-m	allow mixed case, defats to true
	-mixed-case	allow mixed case, defaults to true
	-s	JSON representation of substitution
	-substitutions	JSON representation of substitutions
	-v	display version

EXAMPLE

    slugify "Hello World my friend"


slugify v0.0.11
