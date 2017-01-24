
    This project has merged with https://github.com/rsdoiel/mkpage
    as of January 23, 2017

# mkslides

Converts a Markdown file into a sequence of HTML5 slides.

+ Use Markdown to write your presentation in one file
+ Separate slides by a new line, `--` and another new line (e.g. \n versus \r\n)
+ Apply the simple default template or use your own
+ Control Layout and display with HTML5 and CSS

## Releases and cross compilation

The script [mk-release.sh](./mk-release.sh) cross compiles *mkslides* for Windows, Max OS X, Linux (amd64) and Raspberry Pi (Raspbian/ARM6 and ARM7).
It places all the resulting executable programs in the *dist* folders.

## windows issues

*mkslides* has had very limited Windows testing.  *mkslides.exe* 
presumes the Unix style new line only and not the old DOS/Windows CR/LF type endings.


