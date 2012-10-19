// Copyright (c) 2012 Christopher Rooney
// Code governed by ISC license.  See file LICENSE.

package parurl

import (
	"fmt"
	"net/http"
	"regexp"
)

//Match is the regular expression used to parse parameters from incoming URLs.
//It defaults to integers, optionally preceded by a '/' i.e. \/?([0-9]+) , but
//may be reset.
var Match *regexp.Regexp

const (
	GET = iota
	POST
	PUT
	DELETE
	LAST_METHOD
)

//Handler is the basic type for parametrized method handlers.
type Handler func(http.ResponseWriter, *http.Request, []string)

//URLHandlers is an array of Handlers representing
//GET, POST, PUT and DELETE methods for a base URL.
type URLHandlers [LAST_METHOD]Handler

//BadMethod returns to client a generic Method Not Allowed (405) page. It is used
//to fill in for any missing method handlers.
func BadMethod(w http.ResponseWriter, r *http.Request, _ []string) {
	http.Error(w, fmt.Sprintf("Method %v not permitted for %v.",
		r.Method, r.URL.Path), http.StatusMethodNotAllowed)
}

//NewURLHandlers returns an URLHandlers with either specified or default
//handlers for GET, POST, PUT and DELETE.
func NewURLHandlers(h ...Handler) *URLHandlers {
	p := new(URLHandlers)
	i := copy(p[0:LAST_METHOD], h)
	for ; i < LAST_METHOD; i++ {
		p[i] = BadMethod
	}
	return p
}

//BuildHandler creates an http.HandlerFunc which extracts parameters and
//switches on method.
func BuildHandler(p *URLHandlers) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matches := Match.FindAllString(r.URL.Path, -1)
		if len(matches) == 0 {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case "HEAD":
			w.Header().Set("Content-Length", "0")
		case "GET":
			p[GET](w, r, matches)
		case "POST":
			p[POST](w, r, matches)
		case "PUT":
			p[PUT](w, r, matches)
		case "DELETE":
			p[DELETE](w, r, matches)
		default:
			BadMethod(w, r, matches)
		}
	})
}

func init() {
	Match = regexp.MustCompile("\\/?([0-9]+)")
}
