package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type LogEntry struct {
	Timestamp  string                 `json:"timestamp"`
	Method     string                 `json:"method"`
	Path       string                 `json:"path"`
	Request    map[string]interface{} `json:"request,omitempty"`
	Response   map[string]interface{} `json:"response,omitempty"`
	StatusCode int                    `json:"status_code"`
	Duration   string                 `json:"duration_ms"`
	Error      string                 `json:"error,omitempty"`
	RequestID  string                 `json:"request_id,omitempty"`
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
	body       bytes.Buffer
}

func (w *responseWriter) WriteHeader(code int) {
	if !w.written {
		w.statusCode = code
		w.written = true
		w.ResponseWriter.WriteHeader(code)
	}
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.written {
		w.statusCode = http.StatusOK
		w.written = true
	}
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Capture request body
		var requestBody map[string]interface{}
		if r.Body != nil && r.ContentLength > 0 {
			bodyBytes, _ := io.ReadAll(r.Body)
			json.Unmarshal(bodyBytes, &requestBody)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Wrap response writer to capture response body
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)

		// Parse response body
		var responseBody map[string]interface{}
		if wrapped.body.Len() > 0 {
			json.Unmarshal(wrapped.body.Bytes(), &responseBody)
		}

		logEntry := LogEntry{
			Timestamp:  start.Format(time.RFC3339),
			Method:     r.Method,
			Path:       r.RequestURI,
			Request:    requestBody,
			Response:   responseBody,
			StatusCode: wrapped.statusCode,
			Duration:   duration.String(),
		}

		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			log.Printf("Failed to marshal log entry: %v", err)
			return
		}

		log.Printf("HTTP Request: %s", string(logJSON))
	})
}

var _ func(http.Handler) http.Handler = LoggingMiddleware
