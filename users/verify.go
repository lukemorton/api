package users

import (
	"errors"
)

type VerifyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Verify(users UserFinder, r VerifyRequest) (User, error) {
	err := validateVerifyRequest(r)

	if err != nil {
		return User{}, err
	}

	user, err := users.FindByEmail(r.Email)

	if err != nil {
		return User{}, err
	}

	return user, user.VerifyPassword(r.Password)
}

func validateVerifyRequest(r VerifyRequest) error {
	if r.Email == "" {
		return errors.New("Email address required to verify user")
	} else if r.Password == "" {
		return errors.New("Password required to verify user")
	} else {
		return nil
	}
}
