package handlers

import (
	"net/http"

	"github.com/3timeslazy/chinchilla/storage"
)

// Handlers contains the implementation
// of the application API
type Handlers struct {
	storage storage.Storage
}

// Add sets long url to short
func (h *Handlers) Add(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}

// Redirect redirects from a short url to a long
func (h *Handlers) Redirect(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}

// New returns the new Handlers object
func New(storage storage.Storage) *Handlers {
	return &Handlers{storage: storage}
}
