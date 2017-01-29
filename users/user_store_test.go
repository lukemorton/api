package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	err := users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)
	assertUserStored(t, users)
}

func TestAutoIncrementID(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	users.Create(validUserWithDates("a@gmail.com"))
	users.Create(validUserWithDates("b@gmail.com"))
	assertIncrementedID(t, users)
}

func TestUniqueEmail(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	var err error

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.EqualError(t, err, "Email already taken")
}

func TestCreatePanicOnSQLError(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()
	users.createQuery = "hmm"

	assert.Panics(t, func() {
		users.Create(validUserWithDates("a@gmail.com"))
	})
}

func TestFindByEmail(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()
	users.Create(validUserWithDates("a@gmail.com"))

	user, _ := users.FindByEmail("a@gmail.com")
	assert.Equal(t, "a@gmail.com", user.Email)
}

func TestFindByEmailWithUnknownEmail(t *testing.T) {
	users := ConnectUserStore()
	users.CreateStore()

	_, err := users.FindByEmail("a@gmail.com")
	assert.EqualError(t, err, "Email not recognised")
}

func validUserWithDates(email string) *User {
	return &User{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        email,
		PasswordHash: "bob",
	}
}

func assertUserStored(t *testing.T, db *UserStore) {
	user := User{}
	db.Get(&user, "SELECT email, password_hash FROM users")
	assert.Equal(t, "a@gmail.com", user.Email)
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
