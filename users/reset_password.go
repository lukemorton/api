package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
)

type ResetPasswordUser struct {
	Email string `json:"email"`
}

func ResetPassword(users UserUpdater, r ResetPasswordUser) (string, error) {
	err := validateResetPasswordUser(r)

	if err != nil {
		return "", err
	}

	token := generateToken()
	err = users.UpdateResetTokenHashByEmail(r.Email, tokenHash(token))
	return token, err
}

func validateResetPasswordUser(user ResetPasswordUser) error {
	if user.Email == "" {
		return errors.New("Email address required to reset password")
	} else {
		return nil
	}
}

func generateToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func tokenHash(token string) string {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), 10)

	if err != nil {
		panic(err)
	}

	return string(tokenHash)
}
