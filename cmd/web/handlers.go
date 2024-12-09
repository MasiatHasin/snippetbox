package main

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (h Handler) snippetView(c *gin.Context) {
	id := c.Query("id")
	var err string

	if id != "" {
		id, _ := strconv.Atoi(id)
		snippet, err = h.Repo.Get(id)
	} else {
		//snippet, err = h.Repo.GetAll()
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter 'id' not found"})
		return
	}

	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": snippet})
	}

}

func (h Handler) snippetViewAll(c *gin.Context) {
	var err string

	snippet, err = h.Repo.GetAll()

	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": snippet})
	}

}

func (h Handler) snippetCreate(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	snippet, err := h.Repo.Create(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There has been an error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": snippet})
	}

}

func (h Handler) snippetUpdate(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	id := c.Query("id")
	var err string
	var id_num int

	if id != "" {
		id_num, _ = strconv.Atoi(id)
		snippet, err = h.Repo.Get(id_num)
	} else {
		//snippet, err = h.Repo.GetAll()
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter 'id' not found"})
		return
	}

	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		snippet, _ := h.Repo.Update(body, id_num)
		c.JSON(http.StatusOK, gin.H{"data": snippet})
	}

}

func (h Handler) snippetDelete(c *gin.Context) {
	id := c.Query("id")
	var err string
	var id_num int

	if id != "" {
		id_num, _ = strconv.Atoi(id)
		snippet, err = h.Repo.Get(id_num)
	} else {
		//snippet, err = h.Repo.GetAll()
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter 'id' not found"})
		return
	}

	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		h.Repo.Delete(id_num)
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}

}
