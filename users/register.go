package users

import (
	"errors"
	"time"
)

func Register(db UserCreator, user User) error {
	err := validate(user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err != nil {
		return err
	}

	return db.Create(user)
}

func validate(user User) error {
	if user.Email == "" {
		return errors.New("User requires email address")
	} else {
		return nil
	}
}
