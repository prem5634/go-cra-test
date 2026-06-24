// Package store provides a concurrency-safe, in-memory repository of tasks.
package store

import (
	"errors"
	"sort"
	"sync"
	"time"
)

// ErrNotFound is returned when a requested task does not exist.
var ErrNotFound = errors.New("task not found")

// Task is a single unit of work tracked by the service.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Store is an in-memory, goroutine-safe collection of tasks.
type Store struct {
	mu     sync.RWMutex
	nextID int
	tasks  map[int]Task
}

// New returns an empty Store ready for use.
func New() *Store {
	return &Store{nextID: 1, tasks: make(map[int]Task)}
}

// Create inserts a new task with the given title and returns it.
func (s *Store) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := Task{ID: s.nextID, Title: title, CreatedAt: time.Now().UTC()}
	s.tasks[t.ID] = t
	s.nextID++
	return t
}

// List returns all tasks ordered by ID.
func (s *Store) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Complete marks the task with the given id as done.
func (s *Store) Complete(id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	t.Done = true
	s.tasks[id] = t
	return t, nil
}
