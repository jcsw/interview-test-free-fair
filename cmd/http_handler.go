package main

import (
	http "net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initRouters(mux *http.ServeMux) {
	mux.HandleFunc("/", buildRouters)
	mux.Handle("/metrics", promhttp.Handler())
}

func buildRouters(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			route.handler(w, r)
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ","))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
