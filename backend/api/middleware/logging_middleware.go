package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type LoggingMiddleware struct {
	logging.ILogger
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (rw *responseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}
func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}
func (rw *responseWriter) WriteHeader(statusCode int) {

	if rw.wroteHeader {
		return
	}

	rw.ResponseWriter.WriteHeader(statusCode)
	rw.status = statusCode
	rw.wroteHeader = true

}

func (rw responseWriter) Status() int {
	return rw.status
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

func (l *LoggingMiddleware) Act(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New()
		startTime := time.Now()
		wrappedResponseWriter := wrapResponseWriter(w)

		ctx := context.WithValue(r.Context(), "request_id", requestId)

		l.ILogger.WithGroup("http",
			"method",
			r.Method,
			"path",
			r.URL.Path,
			"query",
			r.URL.RawQuery,
			"user_agent",
			r.UserAgent(),
			"referrer",
			r.Referer(),
			"ip_address",
			r.RemoteAddr,
		)

		handler.ServeHTTP(wrappedResponseWriter, r)

		requestDuration := time.Since(startTime).Milliseconds()
		l.ILogger.With("duration_ms", requestDuration)
		l.ILogger.WithContext(ctx)

		l.ILogger.With("status", wrappedResponseWriter.status)

		if wrappedResponseWriter.Status() >= 400 {
			l.ILogger.Log(logging.ERROR, ctx, "HTTP Request Completed With Error")
			l.ILogger.Clear()
			return
		}

		l.ILogger.Log(logging.INFO, ctx, "HTTP Request Completed")
		l.ILogger.Clear()
	}
}
