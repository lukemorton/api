package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterValidUser(t *testing.T) {
	user := validUser()
	Register(mockUserCreator{}, &user)
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
}

func validUser() User {
	return User{Email: "lukemorton.dev@gmail.com"}
}

type mockUserCreator struct{}

func (users mockUserCreator) Create(user *User) error {
	return nil
}
