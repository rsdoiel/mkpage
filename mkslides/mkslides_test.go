package mkslides

import (
	"strings"
	"testing"
	"text/template"
)

func TestBasic(t *testing.T) {
	src := `

# Opening Slide

## This is an introduction

+ me, <me@example.org>
+ [mygitrepo.org](http://mygitrepo.org)
+ February 31, 2016, Somebody's University, Someplace

--

# Slide Two

| col A | col B | Col C |
|-------|-------|-------|
| A     | One   | 1     |

--

## Slide Three

This is slide three, just a random paragraph of text. Blah, blah, blah, blah, blah, blah, boink!


`

	tmpl, err := template.New("test.tmpl").Parse(DefaultTemplateSource)
	if err != nil {
		t.Errorf("Can't parse DefaultTemplateSource templates %s", err)
		t.FailNow()
	}

	titles := []string{
		"<h1>Opening Slide</h1>",
		"<h1>Slide Two</h1>",
		"<h2>Slide Three</h2>",
	}

	slides := MarkdownToSlides("test.html", "This is just a test", "", "", []byte(src))
	if len(slides) != 3 {
		t.Errorf("Was expected three slides %+v\n", slides)
	}

	for i, slide := range slides {
		s, err := MakeSlideString(tmpl, slide)
		if err != nil {
			t.Errorf("MakeSlideString() failed %d - %s", i, err)
		}
		if strings.Contains(s, titles[i]) == false {
			t.Errorf("Expected %q in slide %d -> %s", titles[i], i, s)
		}
	}
}
