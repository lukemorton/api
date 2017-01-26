package users

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectUserStore() *UserStore {
	db, err := sqlx.Connect("sqlite3", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	return &UserStore{db}
}

type UserStore struct {
	*sqlx.DB
}

func (db *UserStore) CreateStore() {
	db.MustExec(`
		CREATE TABLE users (
			created_at datetime,
			updated_at datetime,
			email text
		);
	`)
}

type UserCreator interface {
	Create(user User) error
}

func (db *UserStore) Create(user User) error {
	q := `
		INSERT INTO users (created_at, updated_at, email)
		VALUES (:created_at, :updated_at, :email)
	`
	_, err := db.NamedExec(q, &user)
	return err
}
