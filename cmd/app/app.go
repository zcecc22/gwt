package main

import (
	"net/http"

	"github.com/zcecc22/gwt"
)

func main() {
	logger := gwt.NewLogger()
	SERVER_LISTEN_ADDRESS := ":10000"
	mux := gwt.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		gwt.RenderJson(w, r, logger, gwt.M{"Hello": "World"}, http.StatusOK)
	})
	mux.HandleFunc("GET /tmpl", func(w http.ResponseWriter, r *http.Request) {
		gwt.RenderTemplate(w, r, logger, "templates/template.tmpl", gwt.M{"text": "hello vincent"}, http.StatusOK)
	})
	gwt.ServeStaticFiles(mux, "GET /static/", "/static/", "static")
	handler := gwt.UseLoggingMiddleware(mux, logger)
	gwt.ListenAndServe(SERVER_LISTEN_ADDRESS, handler, logger)
}
