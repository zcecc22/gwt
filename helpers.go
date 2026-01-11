package gwt

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"text/template"
)

type M map[string]any

func HelperServerError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	logger.Error(err.Error(), "Method:", method, "URI:", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func HelperListenAndServe(addr string, handler http.Handler, logger *slog.Logger) error {
	logger.Info("Starting Server", "Address:", addr)
	return http.ListenAndServe(addr, handler)
}

func HelperNewLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func HelperRenderJson(w http.ResponseWriter, r *http.Request, logger *slog.Logger, data any, statusCode int) {
	js, err := json.Marshal(data)
	if err != nil {
		HelperServerError(w, r, logger, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

func HelperRenderTemplate(w http.ResponseWriter, r *http.Request, logger *slog.Logger, tmplPath string, data any, statusCode int) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		HelperServerError(w, r, logger, err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		HelperServerError(w, r, logger, err)
		return
	}
}

func HelperServeStaticFiles(mux *http.ServeMux, pattern string, urlPath string, dirPath string) {
	mux.Handle(pattern,
		http.StripPrefix(urlPath, http.FileServer(http.Dir(dirPath))))
}
