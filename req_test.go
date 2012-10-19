// Copyright (c) 2012 Christopher Rooney
// Code governed by ISC license.  See file LICENSE.

package parurl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bazHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	fmt.Fprintf(w, "This is the bazHandler.\n it will process %v(%v)\n",
		r.URL.Path, r.URL.Path[5:])
	fmt.Fprintf(w, "We will use the integers: \n")
	for i, v := range matches {
		fmt.Fprintf(w, "%v: %v\n", i, v)
	}
}

var urls = map[string][]Handler{
	"/baz/": {bazHandler, bazHandler},
	"/foo/": {bazHandler},
	"/bar/": {BadMethod, bazHandler, bazHandler},
}

func checkMeth(url string, expected int, t *testing.T,
	meth func(string) (*http.Response, error)) {
	r, err := meth(url)
	if err != nil {
		t.Errorf("%v for %v ", err, "url")
	}
	if r.StatusCode != expected {
		t.Errorf("%v status should be %v but is %v", url, expected, r.StatusCode)
	}
}

func checkGet(url string, expected int, t *testing.T) {
	checkMeth(url, expected, t, http.Get)
}

func checkBoth(url string, expected int, t *testing.T) {
	checkMeth(url, expected, t, http.Head)
	checkMeth(url, expected, t, http.Get)
}

func TestResponses(t *testing.T) {
	for k, v := range urls {
		http.HandleFunc(k, BuildHandler(NewURLHandlers(v...)))
	}
	s := httptest.NewServer(http.DefaultServeMux)
	a := "http://" + s.Listener.Addr().String()
	checkBoth(a+"/baz/234/567", 200, t)
	checkGet(a+"/bar/234/567", 405, t) // no GET handler, but HEAD ok
	checkBoth(a+"/baz/", 404, t)       // no parameterized part
	checkBoth(a+"/dsghsdfg/", 404, t)  // bad url

	r, err := http.Post(a+"/foo/4536345", "image/jpeg", nil)
	if err != nil {
		t.Errorf("%v for %v", err, "foo post")
	}
	if r.StatusCode != 405 {
		t.Errorf("foo post status should be 405 but is %u", r.StatusCode)
	}

	r, err = http.Post(a+"/baz/4536345", "image/jpeg", nil)
	if err != nil {
		t.Errorf("%v for %v", err, "foo post")
	}
	if r.StatusCode != 200 {
		t.Errorf("baz post status should be 200 but is %u", r.StatusCode)
	}
	//should do PUT and DELETE but PITA.  They work.

}
