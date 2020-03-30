package mkpage

import (
	"fmt"
	"io"

	// 3rd Party, Pongo2 library
	"github.com/flosch/pongo2"
)

// MakePongo applies the key/value map to the named template in tmpl and renders to writer and returns an error if something goes wrong
func MakePongo(wr io.Writer, templateName string, keyValues map[string]string) error {
	data, err := ResolveData(keyValues)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	tmpl, err := pongo2.FromFile(templateName)
	if err != nil {
		return fmt.Errorf("Reading Template %q, %s", templateName, err)
	}
	out, err := tmpl.Execute(data)
	if err != nil {
		return fmt.Errorf("Template %q, %s", templateName, err)
	}
	wr.Write([]byte(out))
	return nil
}
