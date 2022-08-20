package main

import (
	fair "interview-test-free-fair/pkg/fair"
	http "net/http"
	regexp "regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

var routes = []route{
	newRoute("GET", "/ping", ping),
	newRoute("GET", "/v1/fairies", fair.SearchHandler),
	newRoute("POST", "/v1/fairies", fair.CreateHandler),
	newRoute("PUT", "/v1/fairies", fair.UpdateHandler),
	newRoute("POST", "/v1/import_data", fair.ImportDataHandler),
}
