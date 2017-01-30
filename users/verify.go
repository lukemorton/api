package users

import (
	"errors"
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

	return user, user.VerifyPassword(v.Password)
}

func validateVerifyUser(user VerifyUser) error {
	if user.Email == "" {
		return errors.New("Email address required to verify user")
	} else if user.Password == "" {
		return errors.New("Password required to verify user")
	} else {
		return nil
	}
}
