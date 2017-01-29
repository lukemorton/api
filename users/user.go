package users

import (
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
