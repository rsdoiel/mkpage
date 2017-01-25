# USAGE

    slugify [OPTIONS] STRING_TO_SLUGIFY

## SYNOPSIS

slugify changes a human readable string into a path or URL friendly
string. E.g. "Hello World" becomes "hello-world"

## OPTIONS

```
	-h	display help
	-l	display license
	-m	allow mixed case, defats to true
	-mixed-case	allow mixed case, defaults to true
	-s	JSON representation of substitution
	-substitutions	JSON representation of substitutions
	-v	display version
```

## EXAMPLE

```
    slugify "Hello World my friend"

```
returns

```
	Hello-World-my-friend
```

