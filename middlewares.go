package gwt

import (
	"fmt"
	"log/slog"
	"net/http"
)

func RequestLogger(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		logger.Info("Received Request", "Address:", ip, "Protocol:", proto, "Method:", method, "URI:", uri)
		next.ServeHTTP(w, r)
	})
}

func RecoverPanic(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			pv := recover()
			if pv != nil {
				w.Header().Set("Connection", "close")
				RenderServerError(w, r, logger, fmt.Errorf("%v", pv))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
