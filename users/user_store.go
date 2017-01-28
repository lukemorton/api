package users

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	defaultCreateQuery = `
		INSERT INTO users (created_at, updated_at, email, password_hash)
		VALUES (:created_at, :updated_at, :email, :password_hash)
	`
)

func ConnectUserStore() *UserStore {
	db, err := sqlx.Connect("sqlite3", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	return &UserStore{db, defaultCreateQuery}
}

type UserStore struct {
	*sqlx.DB
	createQuery string
}

func (db *UserStore) CreateStore() {
	db.MustExec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			email VARCHAR UNIQUE,
			password_hash VARCHAR
		);
	`)
}

type UserCreator interface {
	Create(user *User) error
}

func (db *UserStore) Create(user *User) error {
	result, err := db.NamedExec(db.createQuery, *user)

	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return errors.New("Email already taken")
		} else {
			panic(err)
		}
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	user.Id = id
	return nil
}