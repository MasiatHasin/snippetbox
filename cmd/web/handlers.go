package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"mashiat.snippetbox.test/repository"
)

var snippet interface{}

type Handler struct {
	Repo *repository.DB // Use the repository struct
}

// New creates a new Handler instance
func New(repo *repository.DB) *Handler {
	return &Handler{repo}
}

func (h Handler) snippetView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		id, _ := strconv.Atoi(id)
		snippet = h.Repo.Get(id)
	} else {
		snippet = h.Repo.GetAll()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)

}

func (h Handler) snippetCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	snippet, err = h.Repo.Create(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		message := `{"Error": "There has been an error"}`
		w.Write([]byte(message))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(snippet)
	}

}
