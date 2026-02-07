package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func NewAccessLogMiddleware(logger *zap.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		elapsed := time.Since(start)
		logger.Sugar().Infow(
			"[access]",
			"method", r.Method,
			"url", r.URL,
			"duration", float32(elapsed)/float32(time.Millisecond),
		)
	})
}
