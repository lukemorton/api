package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyValidUser(t *testing.T) {
	user, err := Verify(mockUserFinder{registeredUserForVerification()}, VerifyUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "password",
	})

	assert.Nil(t, err)
	assert.NotNil(t, user.PasswordHash)
}

func TestVerifyUserWithInvalidPassword(t *testing.T) {
	_, err := Verify(mockUserFinder{registeredUserForVerification()}, VerifyUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "not valid",
	})

	assert.EqualError(t, err, "Invalid password")
}

func registeredUserForVerification() User {
	registeredUser, _ := Register(mockUserCreator{}, RegisterUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "password",
	})

	return registeredUser
}

type mockUserFinder struct {
	user User
}

func (users mockUserFinder) FindByEmail(email string) User {
	return users.user
}
