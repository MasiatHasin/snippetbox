package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var snippet interface{}

func home(w http.ResponseWriter, r *http.Request) {
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
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if id != 0 {
		snippet = snippet_model.Get(id)
	} else {
		snippet, _ = snippet_model.GetAll()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)

}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
