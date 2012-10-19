// Copyright (c) 2012 Christopher Rooney
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

/*
Package parurl provides a simple method of parsing and routing parameterized URLs for many http method types, and providing the parameters to handlers.  A typical use is to allow GET, POST, PUT and DELETE calls to the same URL and provide the same URL-derived parameters to each. For example http://example.com/example/123/45/67 would provide 123, 45 and  67, in that order, to each handler for /example/. HEAD requests are handled automatically.

The actual method of determining parameters is a regular expression stored in Match.  By default it is \/?([0-9]+), which will retrieve integers as in the example above, but it may be changed. The entire URL path (that is ,everything after the hostname) will be searches, so the base path should not contain matches for Match.

Typically one writes handlers for various base urls and http methods, creates a URLHandlers from them and calls BuildHandler on the URLHandlers. e.g.

    var urls = map[string][]req.Handler{
         ...
    }
	for k, v := range urls {
		http.HandleFunc(k, parurl.BuildHandler(parurl.NewURLHandlers(v...)))
	}
	http.ListenAndServe(":9090", nil)

A toy example is available in /example, along with a curl session of results.

*/
package parurl
