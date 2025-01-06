package tests

import (
	"satlab-api/internal/articles"
	"satlab-api/internal/database"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestArticleCreateOne(t *testing.T) {
	repo := articles.Repository(database.NewTest())
	article := &articles.Article{
		Title:   "tytul",
		Content: "tu jest tresc",
		Public:  true,
	}
	if err := repo.CreateOne(article); err != nil {
		t.Fatal(err)
	}
}

func TestArticleGetOneById(t *testing.T) {
	db := database.NewTest()
	a := &articles.Article{
		Id:          uuid.NewString(),
		Title:       "tytul",
		Description: "opis",
		Content:     "tu jest tresc",
		Public:      true,
	}
	_, err := db.Exec(`
	INSERT INTO articles (id, title, description, content, created_at, updated_at, public)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`, a.Id, a.Title, a.Description, a.Content, time.Now(), time.Now(), a.Public)
	if err != nil {
		t.Fatal(err)
	}
	repo := articles.Repository(db)
	article, err := repo.GetOneById(a.Id)
	if err != nil {
		t.Fatal(err)
	}
	if article.Title != a.Title {
		t.Fail()
	}
}
