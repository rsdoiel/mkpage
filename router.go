//
// Package provides a light weight package for prefix based routes such as those
// implementing a search API used by cmds/ws/ws.go
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech
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
	"net/http"
	"sort"
	"strings"
)

//
// WSAPI is the web service
//
type WSAPI struct {
	services       map[string]func(w http.ResponseWriter, r *http.Request)
	sortedPrefixes []string
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// AddRoute takes a prefix path string and function with signature func(http.Handler) http.Handler
// adding teh route to the API.
func (wsapi *WSAPI) AddRoute(prefix string, fn func(w http.ResponseWriter, r *http.Request)) error {
	// if route not already registered add route and fn to Services map
	if _, isDefined := wsapi.services[prefix]; isDefined == true {
		return fmt.Errorf("%q already defined", prefix)
	}
	wsapi.services[prefix] = fn

	// Add the prefix to SortedPrefixes and then re-sort by length
	wsapi.sortedPrefixes = append(wsapi.sortedPrefixes, prefix)
	sort.Sort(byLength(wsapi.sortedPrefixes))
	return nil
}

// getPrefix looks up the longest path matching string s and returns
// the matched prefix and true, or empty string and false
func (wsapi *WSAPI) getPrefix(s string) (string, bool) {
	for _, p := range wsapi.sortedPrefixes {
		if len(s) >= len(p) && strings.Contains(s, p) == true {
			return p, true
		}
	}
	return "", false
}

// Router is a simple match prefix route handler for implementing naive route handling
func (wsapi *WSAPI) Router(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If we have a match then call our matched handler with next
		if p, hasPrefix := wsapi.getPrefix(r.URL.Path); hasPrefix == true {
			wsapi.services[p](w, r)
		}
		// Else go straight to the next handler
		next.ServeHTTP(w, r)
	})
}
