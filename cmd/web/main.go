package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mashiat.snippetbox.test/config"
	"mashiat.snippetbox.test/repository"
)

func main() {
	cfg := config.LoadConfig()
	db := repository.Init(cfg)
	repo := repository.New(db)
	h := New(repo)
	r := chi.NewRouter()

	r.Get("/snippet/view", h.snippetView)
	r.Post("/snippet/create", h.snippetCreate)
	fmt.Println("Route registered: POST /snippet/create")
	log.Printf("Starting server on %s", cfg.Port)
	err := http.ListenAndServe(cfg.Port, r)
	log.Fatal(err)
}
