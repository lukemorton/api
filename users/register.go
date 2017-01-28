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

	user := User{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        r.Email,
		PasswordHash: passwordHash(r),
	}

	err = users.Create(&user)
	return user, err
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

func passwordHash(user RegisterUser) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		panic(err)
	}

	return string(passwordHash)
}
