package tests

import (
	"satlab-api/internal/articles"
	"satlab-api/internal/database"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAddImageToArticle(t *testing.T) {
	db := database.NewTest()
	a := &articles.Article{
		Id:          uuid.NewString(),
		Title:       "tytul2",
		Description: "opis2",
		Content:     "tu jest tresc 2",
		Public:      true,
	}
	_, err := db.Exec(`
	INSERT INTO articles (id, title, description, content, created_at, updated_at, public)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`, a.Id, a.Title, a.Description, a.Content, time.Now(), time.Now(), a.Public)
	if err != nil {
		t.Fatal(err)
	}
	repo := articles.Repository(db)
	err = repo.AddImageToArticle(a.Id, "image#1", "path_to_file")
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}
