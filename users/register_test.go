package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterValidUser(t *testing.T) {
	user, err := Register(mockUserCreator{}, RegisterUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "bob",
	})

	assert.Nil(t, err)
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.NotNil(t, user.PasswordHash)
}

func TestRegisterUserWithoutEmail(t *testing.T) {
	_, err := Register(mockUserCreator{}, RegisterUser{
		Password: "bob",
	})

	assert.EqualError(t, err, "User requires email address")
}

func TestRegisterUserWithoutPassword(t *testing.T) {
	_, err := Register(mockUserCreator{}, RegisterUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.EqualError(t, err, "User requires password")
}

func TestUserCreatorError(t *testing.T) {
	users := mockUserCreator{errors.New("Uh oh")}

	_, err := Register(users, RegisterUser{
		Email:    "lukemorton.dev@gmail.com",
		Password: "bob",
	})

	assert.EqualError(t, err, "Uh oh")
}

type mockUserCreator struct {
	err error
}

func (users mockUserCreator) Create(user *User) error {
	return users.err
}
