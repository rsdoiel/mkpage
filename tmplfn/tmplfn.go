/**
 * tmplfn are a collection of functions useful to add to the default
 * Go template/text template definitions. This version
 * of tmplfn is an attempt to converge on Hugo like functions to make
 * the sketchiness of the documentation of using Go templates easier
 * to address by pointing at Hugo's docs.
 *
 * @author R. S. Doiel
 *
 * Copyright (c) 2019, R. S. Doiel
 * All rights not granted herein are expressly reserved by Caltech.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */
package tmplfn

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

var (
	// Version of tmplfn package
	Version = `v0.0.0-rsdoiel`

	// Create a function map for Hugo like functions
	// Reference for functions is https://gohugo.io/functions/
	HugoLike = template.FuncMap{}
)

// Src is a mapping of template source to names and byte arrays.
// It is useful to create a series of defaults templates that can be
// overwritten by user supplied template files.
type Tmpl struct {
	// Holds the function map for templates
	FuncMap template.FuncMap

	// Code holds a map of names to byte arrays, the byte arrays hold the template source code
	// the names can be either filename or other names defined by the implem entor
	Code map[string][]byte
}

// New creates a pointer to a template.Template and  empty map of names to byte arrays
// pointing at an empty byte array
func New(fm template.FuncMap) *Tmpl {
	return &Tmpl{
		FuncMap: fm,
		Code:    map[string][]byte{},
	}
}

// Join take one or more func maps and returns an aggregate one.
func Join(maps ...template.FuncMap) template.FuncMap {
	result := template.FuncMap{}
	for _, m := range maps {
		for key, fn := range m {
			result[key] = fn
		}
	}
	return result
}

// AllFuncs() returns a Join of func maps available in tmplfn
func AllFuncs() template.FuncMap {
	return Join(HugoLike)
}

// ReadFiles takes the given file, filenames or directory name(s) and reads the byte array(s)
// into the Code map.  If a filename is a directory then the directory is scanned
// for files ending in ".tmpl" and those are loaded into the Code map. It does
// NOT parse/assemble templates. The basename in the path is used as the name
// of the template (e.g. templates/page.tmpl would be stored as page.tmpl.
func (t *Tmpl) ReadFiles(fNames ...string) error {
	for _, fname := range fNames {
		if info, err := os.Stat(fname); err != nil {
			return fmt.Errorf("%q, %s", fname, err)
		} else if info.IsDir() == true {
			if files, err := ioutil.ReadDir(fname); err == nil {
				for _, file := range files {
					tname := path.Base(file.Name())
					pname := path.Join(fname, file.Name())
					ext := path.Ext(pname)
					if file.IsDir() != true && ext == ".tmpl" {
						if src, err := ioutil.ReadFile(pname); err != nil {
							return fmt.Errorf("%q, %s", pname, err)
						} else {
							t.Code[tname] = src
						}
					}
				}
			} else {
				return fmt.Errorf("%q, %s", fname, err)
			}
		} else if src, err := ioutil.ReadFile(fname); err == nil {
			tname := path.Base(fname)
			t.Code[tname] = src
		} else {
			return err
		}
	}
	return nil
}

// Add takes a name and source (byte array) and updates t.Code with it.
// It is like Merge but for a single file. The name provided in Add is
// used as the key to the template source code map.
func (t Tmpl) Add(name string, src []byte) error {
	t.Code[name] = src
	if _, ok := t.Code[name]; ok != true {
		return fmt.Errorf("failed to add %s", name)
	}
	return nil
}

// ReadMap works like ReadFiles but takes the name/source pairs from a map rather
// than the file system. It expected template names to end in ".tmpl" like ReadFiles()
// Note the basename of the key provided in the sourceMap is used as the key
// in the Code source code map (e.g. /templates/page.tmpl is stored as page.tmpl)
func (t Tmpl) ReadMap(sourceMap map[string][]byte) error {
	for fname, src := range sourceMap {
		tname := path.Base(fname)
		ext := path.Ext(tname)
		if ext == ".tmpl" {
			t.Code[tname] = src
		}
	}
	if len(t.Code) == 0 {
		return fmt.Errorf("No templates found")
	}
	return nil
}

// Assemble mimics template.ParseFiles() but works with the properties of
// a Tmpl struct.
func (t Tmpl) Assemble() (*template.Template, error) {
	if len(t.Code) == 0 {
		// Mimmic template.ParseFiles() error
		return nil, fmt.Errorf("tmplfn.Assemble(): no template sources to parse")
	}
	var tpl *template.Template
	// Scan the individual templates and parse errors.
	for tName, tSrc := range t.Code {
		s := string(tSrc)
		name := path.Base(tName)
		// This is patterned after template.ParseFiles() internal calls to parseFiles()
		// First template becomes return value if not already defined,
		// use subsequent New calls to associate all the templates together.
		// Otherwise we create a new template associated with t.Template
		var tmpl *template.Template
		if tpl == nil {
			tpl = template.New(name).Funcs(t.FuncMap)
		}
		if name == tpl.Name() {
			tmpl = tpl
		} else {
			tmpl = tpl.New(name).Funcs(t.FuncMap)
		}
		if _, err := tmpl.Parse(s); err != nil {
			return nil, err
		}
	}
	return tpl, nil
}
