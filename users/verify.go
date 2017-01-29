package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type VerifyUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Verify(users UserFinder, v VerifyUser) (User, error) {
	err := validateVerifyUser(v)

	if err != nil {
		return User{}, err
	}

	user, err := users.FindByEmail(v.Email)

	if err != nil {
		return User{}, err
	}

	err = verifyPassword(user.PasswordHash, v.Password)
	return user, err
}

func validateVerifyUser(user VerifyUser) error {
	if user.Email == "" {
		return errors.New("User requires email address")
	} else if user.Password == "" {
		return errors.New("User requires password")
	} else {
		return nil
	}
}

func verifyPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return errors.New("Invalid password")
	}

	return nil
}
