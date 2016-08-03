//
// mkpage is a thought experiment in a light weight template and markdown processor
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights not granted herein are expressly reserved by R. S. Doiel.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package mkpage

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestResolveData(t *testing.T) {
	in := map[string]string{
		"Hello":   "string:Hi there!",
		"Nav":     path.Join("testdata", "nav.md"),
		"Content": path.Join("testdata", "content.md"),
		"Weather": "http://forecast.weather.gov/MapClick.php?lat=13.4712&lon=144.7496&FcstType=json",
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

	in := map[string]string{
		"hello":   "string:Hi there!",
		"nav":     path.Join("testdata", "nav.md"),
		"content": path.Join("testdata", "content.md"),
		"weather": "http://forecast.weather.gov/MapClick.php?lat=9.9667&lon=139.6667&FcstType=json",
	}

	err := MakePage(os.Stdout, src, in, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
