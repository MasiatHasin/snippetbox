package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mashiat.snippetbox.test/config"
)

type Snippet struct {
	ID      int       `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}

type DB struct {
	*gorm.DB
}

// New creates a new DB instance
func New(db *gorm.DB) *DB {
	return &DB{db}
}

func Init(config *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// This will return a specific snippet based on its id.
func (r *DB) Get(id int) *Snippet {
	var snippet Snippet
	if result := r.First(&snippet, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	return &snippet
}

func (r *DB) GetAll() ([]Snippet, error) {
	var snippets []Snippet

	if result := r.Find(&snippets); result.Error != nil {
		fmt.Println(result.Error)
	}

	return snippets, nil
}
