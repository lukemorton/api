package users

import (
	"errors"
	"time"
)

func Register(users UserCreator, user *User) error {
	err := validate(*user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err != nil {
		return err
	}

	return users.Create(user)
}

func validate(user User) error {
	if user.Email == "" {
		return errors.New("User requires email address")
	} else {
		return nil
	}
}
