parurl is a very simple, very small, URL request router and parametrizer, for when a whole framework is overkill. parurl supports HEAD, GET, POST, PUT, and DELETE for RESTful applications. It uses http.DefaultServeMux so other request types may be mixed in using standard calls like http.HandleFunc, etc.  

Pretty much it lets you easily receive a request like http://example.com/example/123/98/45, route it to the correct handler for GET /example/, POST /example/ etc. and supply the handler with {123, 98, 45}. Integer parameters are the default, but can be overridden.

This was meant to be my first nontrivial go project, but it turned out to be dead easy, and well, trivial.

####Installation
go get github.com/crooney/parurl

####Example
go get github.com/crooney/parurl/example

####Dependencies
None

####Documentation
doc.go to read here or godoc once installed

####How Small Is It?
54 sloc