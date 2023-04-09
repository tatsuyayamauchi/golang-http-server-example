package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct{}

type RequestLog struct {
	URL           string        `json:"url"`
	RemoteAddress string        `json:"remote_address"`
	Method        string        `json:"method"`
	Latency       time.Duration `json:"latency"`
	RequestedAt   time.Time     `json:"requested_at"`
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		var buf bytes.Buffer
		req := &RequestLog{
			URL:           r.URL.String(),
			Method:        r.Method,
			RemoteAddress: r.RemoteAddr,
			RequestedAt:   start,
			Latency:       time.Duration(time.Now().Sub(start).Microseconds()),
		}

		if err := json.NewEncoder(&buf).Encode(req); err != nil {
			log.Printf("encode err: %v\n", err)
		} else {
			log.Println(buf.String())
		}
	})
}
