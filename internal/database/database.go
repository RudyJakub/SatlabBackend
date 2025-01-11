package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}

	if err := initTables(db); err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func NewTest() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	if err := initTables(db); err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func initTables(db *sql.DB) error {
	var err error
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS articles(
		id text,
		title text,
		description text,
		content text,
		created_at integer,
		updated_at integer,
		public integer,
		CONSTRAINT articles_pkey PRIMARY KEY (id)
	)`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS images(
		id text,
		title text,
		uploaded_at integer,
		image_location text,
		article_id text not null,
		CONSTRAINT images_pkey PRIMARY KEY (id),
		FOREIGN KEY (article_id) REFERENCES articles(id)
	)`)

	return err
}
