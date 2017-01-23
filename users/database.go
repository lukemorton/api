package users

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDB() *DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	db.MustExec(`
    CREATE TABLE users (
      created_at datetime,
      updated_at datetime,
      email text
    );
  `)

	if err != nil {
		log.Fatalln(err)
	}

	return &DB{db}
}

type DB struct {
	*sqlx.DB
}

type UserCreator interface {
	Create(user User) error
}

func (db *DB) Create(user User) error {
	q := `
		INSERT INTO users (created_at, updated_at, email)
		VALUES (:created_at, :updated_at, :email)
	`
	_, err := db.NamedExec(q, &user)
	return err
}
