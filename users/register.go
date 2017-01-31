package users

import (
	"errors"
)

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(users UserCreator, r RegisterUser) (User, error) {
	err := validateRegisterUser(r)

	if err != nil {
		return User{}, err
	}

	user := User{Email: r.Email}
	user.SetPassword(r.Password)

	err = users.Create(&user)
	return user, err
}

func validateRegisterUser(user RegisterUser) error {
	if user.Email == "" {
		return errors.New("User requires email address")
	} else if user.Password == "" {
		return errors.New("User requires password")
	} else {
		return nil
	}
}
