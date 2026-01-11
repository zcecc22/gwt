package gwt

import (
	"log/slog"
	"net/http"
)

func RouterNewServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func RouterUseLoggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return MiddlewareRequestLogger(MiddlewareRecoverPanic(next, logger), logger)
}
