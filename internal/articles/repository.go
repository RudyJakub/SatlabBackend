package articles

import (
	"satlab-api/internal/database"
	"time"

	"github.com/google/uuid"
)

type repo struct {
	db database.Service
}

func Repository(db database.Service) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateOne(a *Article) error {
	_, err := r.db.Exec(`
	INSERT INTO articles (id, title, description, content, created_at, updated_at, public)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`, uuid.NewString(), a.Title, a.Description, a.Content, time.Now(), time.Now(), a.Public)
	return err
}

func (r *repo) DeleteOne(id string) error {
	_, err := r.db.Exec(`DELETE FROM articles WHERE id=$1`, id)
	return err
}

func (r *repo) UpdateOne(id string, a *Article) error {
	_, err := r.db.Exec(`UPDATE articles SET
	title=$1, description=$2, content=$3, updated_at=$4, public=$5
	WHERE id=$6`, a.Title, a.Description, a.Content, time.Now(), a.Public, id)
	return err
}

func (r *repo) GetOneById(id string) (*Article, error) {
	var article Article
	err := r.db.QueryRow(`SELECT * FROM articles WHERE id=$1;`, id).Scan(
		&article.Id,
		&article.Title,
		&article.Description,
		&article.Content,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.Public,
	)
	return &article, err
}

func (r *repo) GetAll() ([]Article, error) {
	var articles []Article
	rows, err := r.db.Query(`SELECT * FROM articles ORDER BY updated_at DESC;`)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *repo) GetAllPublic() ([]Article, error) {
	var articles []Article
	rows, err := r.db.Query(`SELECT * FROM articles WHERE public=1 ORDER BY updated_at DESC;`)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *repo) AddImageToArticle(article_id string, title string, path string) error {
	_, err := r.db.Query(`INSERT INTO images (id, title, uploaded_at, image_location, article_id)
	VALUES ($1, $2, $3, $4, $5)`, uuid.NewString(), title, time.Now(), path, article_id)

	return err
}
