package main

import (
	"log"
	"net/http"

	"mashiat.snippetbox.test/repository"
)

var snippet_model *repository.DBModel

func main() {
	config := loadConfig()

	snippet_model = repository.ConnectDB(config.DB_URL)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/jsonview", jsonview)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", config.Port)
	err := http.ListenAndServe(config.Port, mux)
	log.Fatal(err)
}
