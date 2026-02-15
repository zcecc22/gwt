package gwt

import (
	"log/slog"
	"net/http"
)

func NewServeMux() *http.ServeMux {
	return http.NewServeMux()
}

func UseLoggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return RequestLogger(RecoverPanic(next, logger), logger)
}
