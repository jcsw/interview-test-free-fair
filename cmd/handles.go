package main

import (
	context "context"
	sys "interview-test-free-fair/pkg/infra/system"
	http "net/http"
	"strconv"
	time "time"

	uuid "github.com/google/uuid"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func wrapResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func tracing() func(http.Handler) http.Handler {
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

			httpRequestMetric(r.URL.Path, r.Method, strconv.Itoa(wrw.statusCode), float64(time.Since(start).Milliseconds()))
			sys.Info("[tracing][%s][%s][%s][%d][%v]", requestID, r.Method, r.URL.Path, wrw.statusCode, time.Since(start))
		})
	}
}
