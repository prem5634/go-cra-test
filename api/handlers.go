// Package api exposes the HTTP handlers for the task service.
package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/prembito/go-cra-test/auth"
	"gitlab.com/prembito/go-cra-test/store"
)

// Handler bundles the dependencies shared by the HTTP routes.
type Handler struct {
	Store  *store.Store
	Issuer *auth.Issuer
}

type createRequest struct {
	Title string `json:"title"`
}

// IssueToken handles POST /token and returns a signed JWT for ?user=.
func (h *Handler) IssueToken(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		writeError(w, http.StatusBadRequest, "user is required")
		return
	}
	tok, err := h.Issuer.Token(user)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not issue token")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"token": tok})
}

// CreateTask handles POST /tasks.
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.Title) == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}
	writeJSON(w, http.StatusCreated, h.Store.Create(req.Title))
}

// ListTasks handles GET /tasks.
func (h *Handler) ListTasks(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.Store.List())
}

// CompleteTask handles POST /tasks/{id}/complete.
func (h *Handler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}
	t, err := h.Store.Complete(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}
