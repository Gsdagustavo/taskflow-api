package router

import (
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		slog.InfoContext(ctx, "request received",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("host", r.Host),
		)

		next.ServeHTTP(w, r)
	})
}
