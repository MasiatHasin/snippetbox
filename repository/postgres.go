package repository

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetForm struct {
	Title   string
	Content string
	Expires int
}

type DBModel struct {
	db *pgxpool.Pool
}

func ConnectDB(db_url string) *DBModel {
	db, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	return &DBModel{db: db}
}

// This will return a specific snippet based on its id.
func (repo *DBModel) Get(id int) *Snippet {
	var snippet Snippet
	err := repo.db.QueryRow(context.Background(), "SELECT * FROM admin_app_snippet WHERE id=$1", id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	return &snippet
}

func (repo *DBModel) GetAll() ([]Snippet, error) {
	var snippets []Snippet

	rows, err := repo.db.Query(context.Background(), "SELECT * FROM admin_app_snippet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var snippet Snippet
		if err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires); err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
