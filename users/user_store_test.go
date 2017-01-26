package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	err := users.Create(validUserWithDates())
	assert.Nil(t, err)
	assertUserStored(t, users)
}

func TestAutoIncrementID(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	users.Create(validUserWithDates())
	users.Create(validUserWithDates())
	assertIncrementedID(t, users)
}

func validUserWithDates() User {
	return User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "lukemorton.dev@gmail.com",
	}
}

func assertUserStored(t *testing.T, db *UserStore) {
	var email string
	err := db.Get(&email, "SELECT email FROM users")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, email, "lukemorton.dev@gmail.com")
}

func assertIncrementedID(t *testing.T, db *UserStore) {
	var id []int
	err := db.Select(&id, "SELECT id FROM users")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(id), 2)
	assert.Equal(t, id[0], 1)
	assert.Equal(t, id[1], 2)
}
