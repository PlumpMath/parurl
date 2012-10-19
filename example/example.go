// Copyright (c) 2012 Christopher Rooney
// Code governed by ISC license.  See file LICENSE.

package main

import (
	"fmt"
	"github.com/crooney/parurl"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request, matches []string) {
	fmt.Fprintf(w, "We will use the integers: \n")
	for i, v := range matches {
		fmt.Fprintf(w, "%v: %v\n", i, v)
	}
}

var urls = map[string][]parurl.Handler{
	"/baz/": {handler, handler},                            //GET, POST
	"/foo/": {handler},                                     //GET
	"/bar/": {parurl.BadMethod, handler, handler, handler}, //POST, PUT, DELETE
}

func main() {
	for k, v := range urls {
		http.HandleFunc(k, parurl.BuildHandler(parurl.NewURLHandlers(v...)))
	}
	http.ListenAndServe(":9090", nil)
}

/* Sample session
[0 ~]$ curl -XGET localhost:9090/baz/34/56
We will use the integers:
0: /34
1: /56
[0 ~]$ curl -XPOST localhost:9090/baz/34/56
We will use the integers:
0: /34
1: /56
[0 ~]$ curl -XDELETE localhost:9090/baz/34/56
Method DELETE not permitted for /baz/34/56.
[0 ~]$ curl -XDELETE localhost:9090/bar/34/56
We will use the integers:
0: /34
1: /56
[0 ~]$ curl -XGET localhost:9090/bar/34/56
Method GET not permitted for /bar/34/56.
*/
