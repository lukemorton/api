package users

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strings"
	"time"
)

const (
	defaultCreateTableQuery = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER AUTO_INCREMENT PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			email VARCHAR(256),
			password_hash VARCHAR(128),
			reset_token_hash VARCHAR(128),
			UNIQUE(email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`
	defaultTruncateTableQuery = `
		TRUNCATE TABLE users;
	`
	defaultDropTableQuery = `
		DROP TABLES users;
	`
	defaultCreateQuery = `
		INSERT INTO users
			(created_at, updated_at, email, password_hash, reset_token_hash)
	 	VALUES
			(:created_at, :updated_at, :email, :password_hash, "")
	`
	defaultUpdateResetTokenHashByIdQuery = `
		UPDATE users SET reset_token_hash = ?, updated_at = ? WHERE id = ?
	`
	defaultUpdatePasswordHashByIdQuery = `
		UPDATE users SET password_hash = ?, updated_at = ? WHERE id = ?
	`
	defaultFindByEmailQuery = `
		SELECT * FROM users WHERE email = ?
	`
)

func SQLUserStore() *sqlUserStore {
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		log.Fatalln("No DATABASE_URL defined")
	}

	db, err := sqlx.Connect("mysql", databaseUrl)

	if err != nil {
		log.Fatalln(err)
	}

	return &sqlUserStore{
		db,
		defaultCreateQuery,
		defaultUpdateResetTokenHashByIdQuery,
		defaultUpdatePasswordHashByIdQuery,
		defaultFindByEmailQuery,
	}
}

type sqlUserStore struct {
	*sqlx.DB
	createQuery                   string
	updateResetTokenHashByIdQuery string
	updatePasswordHashByIdQuery   string
	findByEmailQuery              string
}

func (db *sqlUserStore) CreateStore() {
	db.MustExec(defaultCreateTableQuery)
}

func (db *sqlUserStore) ClearStore() {
	db.MustExec(defaultTruncateTableQuery)
}

func (db *sqlUserStore) DeleteStore() {
	db.MustExec(defaultDropTableQuery)
}

func (db *sqlUserStore) Create(user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result, err := db.NamedExec(db.createQuery, *user)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
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

func (db *sqlUserStore) UpdateResetTokenHash(user *User) {
	db.updateField(user, db.updateResetTokenHashByIdQuery, user.ResetTokenHash)
}

func (db *sqlUserStore) UpdatePasswordHash(user *User) {
	db.updateField(user, db.updatePasswordHashByIdQuery, user.PasswordHash)
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

func (db *sqlUserStore) updateField(user *User, query string, value interface{}) {
	user.UpdatedAt = time.Now()
	result := db.MustExec(query, value, user.UpdatedAt, user.Id)
	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	if rows == 0 {
		panic("No rows updated")
	}
}
