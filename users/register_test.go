package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterValidUser(t *testing.T) {
	user, err := Register(mockUserCreator{}, validUser())
	assert.Nil(t, err)
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.NotNil(t, user.PasswordHash)
}

func validUser() RegisterUser {
	return RegisterUser{
		Email: "lukemorton.dev@gmail.com",
		Password: "bob",
	}
}

type mockUserCreator struct{}

func (users mockUserCreator) Create(user *User) error {
	return nil
}
