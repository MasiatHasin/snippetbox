package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Snippet struct {
	ID      int       `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}

type SnippetForm struct {
	Title   string
	Content string
	Expires int
}

type DBModel struct {
	db *gorm.DB
}

func ConnectDB(db_url string) *DBModel {
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Snippet{})

	return &DBModel{db: db}
}

// This will return a specific snippet based on its id.
func (repo *DBModel) Get(id int) *Snippet {
	var snippet Snippet
	if result := repo.db.First(&snippet, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Print(&snippet)
	return &snippet
}

func (repo *DBModel) GetAll() ([]Snippet, error) {
	var snippets []Snippet

	if result := repo.db.Find(&snippets); result.Error != nil {
		fmt.Println(result.Error)
	}

	return snippets, nil
}
