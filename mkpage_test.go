package mkpage

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestResolveData(t *testing.T) {
	in := map[string][]byte{
		"Hello":   []byte("string:Hi there!"),
		"Nav":     []byte(path.Join("testdata", "nav.md")),
		"Content": []byte(path.Join("testdata", "content.md")),
		"Weather": []byte("http://forecast.weather.gov/MapClick.php?lat=38.4247341&lon=-86.9624086&FcstType=json"),
	}
	data, err := ResolveData(in, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", data)
}

func TestMakePage(t *testing.T) {
	src := `
Hello {{.hello}}

Nav: {{.nav}}

Content: {{.content}}

Weather: {{.weather}}
`

	in := map[string][]byte{
		"hello":   []byte("string:Hi there!"),
		"nav":     []byte(path.Join("testdata", "nav.md")),
		"content": []byte(path.Join("testdata", "content.md")),
		"weather": []byte("http://forecast.weather.gov/MapClick.php?lat=9.9667&lon=139.6667&FcstType=json"),
	}

	err := MakePage(os.Stdout, src, Text, in, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
