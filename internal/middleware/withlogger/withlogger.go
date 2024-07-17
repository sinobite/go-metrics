package withlogger

import (
	"github.com/sinobite/go-metrics/internal/logger"
	"net/http"
	"time"
)

type MiddlewareFunc func(next http.Handler) http.Handler

func New(log logger.Logger) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()

			responseData := &responseData{
				status: 0,
				size:   0,
			}
			lw := loggingResponseWriter{
				ResponseWriter: writer,
				responseData:   responseData,
			}

			next.ServeHTTP(&lw, request)

			duration := time.Since(start)
			log.ZapLog.Infow("Request data", "uri", request.RequestURI, "method", request.Method, "duration", duration)

			log.ZapLog.Infow("Response data", "status", responseData.status, "size", responseData.size)

		})
	}
}

// Оставил здесь, пока не уверен что это нужно переиспользовать
type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}
