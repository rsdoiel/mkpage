
# mkslides

Converts a Markdown file into a sequence of HTML5 slides.

+ Use Markdown to write your presentation in one file
+ Separate slides by a new line, `--` and another new line (e.g. \n versus \r\n)
+ Apply the simple default template or use your own
+ Control Layout and display with HTML5 and CSS


## Example

A simple markdown producing three slides (assumed you named your file
three-slides.md)

```markdown
    # Title Page

    ## By M.E. Person

    --

    # Second Page slide

    + one
        + subone
        + subtwo
    + two
    + three

    --

    # Finale Page

    This is the end.
```

To make a set of slides

```shell
    mkslides three-slides.md
```

Output should look something like

```shell
    Wrote 01-three-slides.html
    Wrote 02-three-slides.html
    Wrote toc-three-slides.html
```

Now you should see a set of slides for your presentation.

## windows issues

*mkslides* has had very limited Windows testing.  *mkslides.exe* 
presumes the Unix style new line only and not the old DOS/Windows CR/LF type endings.


