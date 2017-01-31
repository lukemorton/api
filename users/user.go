package users

import (
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"          db:"password_hash"`
	ResetTokenHash string    `json:"-"          db:"reset_token_hash"`
}

func (user *User) SetPassword(password string) {
	user.PasswordHash = hash(password)
}

func (user *User) VerifyPassword(password string) error {
	if verifyHash(user.PasswordHash, password) {
		return nil
	} else {
		return errors.New("Invalid password")
	}
}

func (user *User) GenerateResetToken() string {
	token := generateToken()
	user.ResetTokenHash = hash(token)
	return token
}

func hash(s string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(s), 10)

	if err != nil {
		panic(err)
	}

	return string(h)
}

func verifyHash(h string, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(s))
	return err == nil
}

func generateToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
