package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"mashiat.snippetbox.test/config"
	"mashiat.snippetbox.test/repository"
)

func main() {
	cfg := config.LoadConfig()
	db := repository.Init(cfg)
	repo := repository.New(db)
	h := New(repo)
	r := gin.Default()

	r.GET("/snippet/view", h.snippetView)
	r.GET("/snippet/view/all", h.snippetViewAll)
	r.POST("/snippet/create", h.snippetCreate)
	r.POST("/snippet/update", h.snippetUpdate)
	r.POST("/snippet/delete", h.snippetDelete)
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	})
	log.Printf("Starting server on %s", cfg.Port)
	err := http.ListenAndServe(cfg.Port, r)
	log.Fatal(err)
}
