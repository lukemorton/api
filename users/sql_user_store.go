package users

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	defaultCreateTableQuery = `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			email VARCHAR UNIQUE,
			password_hash VARCHAR
		);
	`
	defaultCreateQuery = `
		INSERT INTO users (created_at, updated_at, email, password_hash)
		VALUES (:created_at, :updated_at, :email, :password_hash)
	`
	defaultFindByEmailQuery = `
		SELECT * FROM users WHERE email = ?
	`
)

func SQLUserStore() *sqlUserStore {
	db, err := sqlx.Connect("sqlite3", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	return &sqlUserStore{
		db,
		defaultCreateQuery,
		defaultFindByEmailQuery,
	}
}

type sqlUserStore struct {
	*sqlx.DB
	createQuery string
	findByEmailQuery string
}

func (db *sqlUserStore) CreateStore() {
	db.MustExec(defaultCreateTableQuery)
}

func (db *sqlUserStore) Create(user *User) error {
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

func (db *sqlUserStore) FindByEmail(email string) (User, error) {
	user := User{}
	err := db.Get(&user, db.findByEmailQuery, email)

	if err != nil {
		return user, errors.New("Email not recognised")
	}

	return user, nil
}
