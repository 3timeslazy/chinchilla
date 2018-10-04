package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/3timeslazy/chinchilla/handlers"
	"github.com/3timeslazy/chinchilla/storage"
	"github.com/3timeslazy/chinchilla/storage/postgres"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func run() error {
	h := handlers.New(mustStorage())
	r := mux.NewRouter()

	r.HandleFunc("/add", h.Add).Methods("POST")
	r.HandleFunc("/{short:[A-Za-z0-9]+}", h.Redirect).Methods("GET")

	http.Handle("/", r)
	fmt.Println("listen at :8080")
	return http.ListenAndServe(":8080", muxhandlers.LoggingHandler(os.Stderr, r))
}

func mustStorage() storage.Storage {
	db, err := sql.Open("postgres", "postgres://chinchilla:chinchilla@localhost:5432/chinchilla?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return postgres.New(db)
}
