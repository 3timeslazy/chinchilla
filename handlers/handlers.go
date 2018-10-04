package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/3timeslazy/chinchilla/storage"
)

// Handlers contains the implementation
// of the application API
type Handlers struct {
	storage storage.Storage
}

type reqsPostAdd struct {
	URL string `json:"url"`
}

// Add sets long url to short
func (h *Handlers) Add(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}

// Redirect redirects from a short url to a long
func (h *Handlers) Redirect(w http.ResponseWriter, r *http.Request) {
	short := mux.Vars(r)["short"]

	long, err := h.storage.Extract(short)
	switch err {
	case nil:
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error. pls try later"))
		return
	}

	http.Redirect(w, r, "http://"+long, http.StatusFound)
}

// New returns the new Handlers object
func New(storage storage.Storage) *Handlers {
	return &Handlers{storage: storage}
}
