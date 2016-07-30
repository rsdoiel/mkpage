package mkpage

import (
	"bytes"
	"fmt"
	HTMLTemplate "html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	TextTemplate "text/template"

	// 3rd Party Packages
	"github.com/russross/blackfriday"
)

const (
	Version = "v0.0.1"
	HTML    = iota
	Text    = iota
)

// ResolveData takes a data map and reads in the files and URL sources as needed turning
// the data into strings to be applied to the template.
func ResolveData(data map[string][]byte, useMarkdownProcessor bool) (map[string]string, error) {
	var out map[string]string

	out = make(map[string]string)
	for key, val := range data {
		switch {
		case bytes.HasPrefix(val, []byte("string:")) == true:
			out[key] = string(bytes.TrimPrefix(val, []byte("string:")))
		case bytes.HasPrefix(val, []byte("http://")) == true || bytes.HasPrefix(val, []byte("https://")):
			resp, err := http.Get(string(val))
			if err != nil {
				return out, err
			}
			defer resp.Body.Close()
			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return out, err
			}
			out[key] = string(buf)
		default:
			buf, err := ioutil.ReadFile(string(val))
			if err != nil {
				return out, err
			}
			ext := path.Ext(string(val))
			fmt.Printf("DEBUG ext: %s\n", ext)
			if useMarkdownProcessor == true {
				out[key] = string(blackfriday.MarkdownCommon(buf))
			} else {
				out[key] = string(buf)
			}
		}
	}
	return out, nil
}

// MakePage applies the provided data to the template provided and renders to writer and returns an error if something goes wrong
func MakePage(wr io.Writer, templateSource string, templateType int, keyValues map[string][]byte, useMarkdownProcessor bool) error {
	data, err := ResolveData(keyValues, useMarkdownProcessor)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	switch templateType {
	case HTML:
		tmpl, err := HTMLTemplate.New("html").Parse(templateSource)
		if err != nil {
			return err
		}
		return tmpl.Execute(wr, data)
	default:
		tmpl, err := TextTemplate.New("text").Parse(templateSource)
		if err != nil {
			return err
		}
		return tmpl.Execute(wr, data)
	}
}
