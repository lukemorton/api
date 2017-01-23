package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	user := validUserWithDates()
	db := ConnectDB()

	err := db.Create(user)
	assert.Nil(t, err)
}

func validUserWithDates() User {
	return User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "lukemorton.dev@gmail.com",
	}
}
