package http

import (
	context "context"
	http "net/http"
	"strconv"
	time "time"

	sys "interview-test-free-fair/pkg/sys"

	uuid "github.com/google/uuid"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func wrapResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func Tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			ctx := context.WithValue(r.Context(), sys.ContextKeyRequestID, requestID)
			w.Header().Set("X-Request-Id", requestID)

			wrw := wrapResponseWriter(w)
			next.ServeHTTP(wrw, r.WithContext(ctx))

			sys.HttpRequestMetric(r.URL.Path, r.Method, strconv.Itoa(wrw.statusCode), float64(time.Since(start).Milliseconds()))
			sys.LogInfo("[tracing][%s][%s][%s][%d][%v]", requestID, r.Method, r.URL.RequestURI(), wrw.statusCode, time.Since(start))
		})
	}
}
