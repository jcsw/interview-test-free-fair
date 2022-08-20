package http

import (
	http "net/http"
	"strings"

	app "interview-test-free-fair/pkg/app"
	mariadb "interview-test-free-fair/pkg/mariadb"
)

func BuildHandlers(mux *http.ServeMux) {

	fairService := app.FairServiceMariaDb{BD: mariadb.RetrieveClient()}
	fairHandler := &app.FairHandler{Service: &fairService}

	routes = append(routes, newRoute("GET", "/v1/fairies/search", fairHandler.Search))
	routes = append(routes, newRoute("GET", "/v1/fairies", fairHandler.Find))
	routes = append(routes, newRoute("POST", "/v1/fairies", fairHandler.Create))
	routes = append(routes, newRoute("PUT", "/v1/fairies", fairHandler.Update))
	routes = append(routes, newRoute("DELETE", "/v1/fairies", fairHandler.Delete))
	routes = append(routes, newRoute("POST", "/v1/import_data", fairHandler.ImportData))

	mux.HandleFunc("/", buildRouters)
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
