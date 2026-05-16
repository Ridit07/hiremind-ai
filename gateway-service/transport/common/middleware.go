package http

import (
	"bytes"
	"context"
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

		var requestBody map[string]interface{}
		if r.Body != nil && r.ContentLength > 0 {
			bodyBytes, _ := io.ReadAll(r.Body)
			json.Unmarshal(bodyBytes, &requestBody)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)

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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "missing authorization header",
			})
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "invalid authorization header format",
			})
			return
		}

		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "missing X-User-ID header",
			})
			return
		}

		log.Printf("Auth middleware: user_id=%s, auth_header_present=true", userID)

		token := authHeader[7:]

		ctx := context.WithValue(r.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
