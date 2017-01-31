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
			password_hash VARCHAR,
			reset_token_hash VARCHAR NOT NULL
		);
	`
	defaultCreateQuery = `
		INSERT INTO users
			(created_at, updated_at, email, password_hash, reset_token_hash)
	 	VALUES
			(:created_at, :updated_at, :email, :password_hash, "")
	`
	defaultUpdateResetTokenHashByIdQuery = `
		UPDATE users SET reset_token_hash = ? WHERE id = ?
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
		defaultUpdateResetTokenHashByIdQuery,
		defaultFindByEmailQuery,
	}
}

type sqlUserStore struct {
	*sqlx.DB
	createQuery                   string
	updateResetTokenHashByIdQuery string
	findByEmailQuery               string
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

func (db *sqlUserStore) UpdateResetTokenHash(user *User) error {
	result := db.MustExec(db.updateResetTokenHashByIdQuery, user.ResetTokenHash, user.Id)
	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	if rows == 0 {
		panic(user)
	}

	return nil
}

func (db *sqlUserStore) FindByEmail(email string) (User, error) {
	user := User{}
	err := db.Get(&user, db.findByEmailQuery, email)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return user, errors.New("Email not recognised")
		} else {
			panic(err)
		}
	}

	return user, nil
}
