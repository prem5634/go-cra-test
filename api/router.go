// Package api wires routes and middleware onto an HTTP mux.
package api

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gitlab.com/prembito/go-cra-test/auth"
)

// NewRouter builds the service mux with all routes registered. Mutating
// routes are guarded by bearer-token authentication.
func NewRouter(h *Handler, logRequests bool) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /token", h.IssueToken)
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", requireAuth(h.CreateTask))
	mux.HandleFunc("POST /tasks/{id}/complete", requireAuth(h.CompleteTask))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	if !logRequests {
		return mux
	}
	return logging(mux)
}

// requireAuth rejects requests that lack a valid bearer token.
func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if _, err := auth.Verify(raw); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// logging records the method, path, and latency of each request.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
