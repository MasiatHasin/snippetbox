package main

import (
	"encoding/json"
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

/* func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func jsonview(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	rawJSON := `{"status":"success", "message":"This is a raw JSON string response"}`
	w.Write([]byte(rawJSON))
} */

func (h Handler) snippetView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		id, _ := strconv.Atoi(id)
		snippet = h.Repo.Get(id)
	} else {
		snippet, _ = h.Repo.GetAll()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)

}

/* func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
} */
