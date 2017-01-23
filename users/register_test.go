package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterValidUser(t *testing.T) {
	user := validUser()
	Register(mockDB{}, user)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
}

func validUser() User {
	return User{Email: "lukemorton.dev@gmail.com"}
}

type mockDB struct{}

func (db mockDB) Create(user User) error {
	return nil
}
