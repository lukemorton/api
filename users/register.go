package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(users UserCreator, r RegisterUser) (User, error) {
	err := validate(r)

	if err != nil {
		return User{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)

	if err != nil {
		return User{}, err
	}

	user := User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email: r.Email,
		PasswordHash: string(passwordHash),
	}

	err = users.Create(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func validate(user RegisterUser) error {
	if user.Email == "" {
		return errors.New("User requires email address")
	} else if user.Password == "" {
		return errors.New("User requires password")
	} else {
		return nil
	}
}
