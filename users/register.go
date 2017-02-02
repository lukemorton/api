package users

import (
	"errors"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(users UserCreator, r RegisterRequest) (User, error) {
	err := validateRegisterRequest(r)

	if err != nil {
		return User{}, err
	}

	user := User{Email: r.Email}
	user.SetPassword(r.Password)

	err = users.Create(&user)
	return user, err
}

func validateRegisterRequest(r RegisterRequest) error {
	if r.Email == "" {
		return errors.New("User requires email address")
	} else if r.Password == "" {
		return errors.New("User requires password")
	} else {
		return nil
	}
}
