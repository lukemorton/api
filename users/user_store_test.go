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

func validUserWithDates() *User {
	return &User{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        "lukemorton.dev@gmail.com",
		PasswordHash: "bob",
	}
}

func assertUserStored(t *testing.T, db *UserStore) {
	user := User{}
	db.Get(&user, "SELECT email, password_hash FROM users")
	assert.Equal(t, "lukemorton.dev@gmail.com", user.Email)
	assert.Equal(t, "bob", user.PasswordHash)
}

func assertIncrementedID(t *testing.T, db *UserStore) {
	var id []int
	err := db.Select(&id, "SELECT id FROM users")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(id))
	assert.Equal(t, 1, id[0])
	assert.Equal(t, 2, id[1])
}
