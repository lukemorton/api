package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyValidUser(t *testing.T) {
	user, err := Verify(mockUserFinder{user: registeredUserForVerification()}, VerifyUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "password",
	})

	assert.Nil(t, err)
	assert.NotNil(t, user.PasswordHash)
}

func TestVerifyUserWithoutEmail(t *testing.T) {
	_, err := Verify(mockUserFinder{}, VerifyUser{
		Password: "bob",
	})

	assert.EqualError(t, err, "Email address required to verify user")
}

func TestVerifyUserWithoutPassword(t *testing.T) {
	_, err := Verify(mockUserFinder{}, VerifyUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.EqualError(t, err, "Password required to verify user")
}

func TestVerifyUserWithInvalidEmail(t *testing.T) {
	_, err := Verify(mockUserFinder{err: errors.New("Not found")}, VerifyUser{
		Email:    "notfound@gmail.com",
		Password: "password",
	})

	assert.EqualError(t, err, "Not found")
}

func TestVerifyUserWithInvalidPassword(t *testing.T) {
	_, err := Verify(mockUserFinder{user: registeredUserForVerification()}, VerifyUser{
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
	err error
}

func (users mockUserFinder) FindByEmail(email string) (User, error) {
	return users.user, users.err
}
