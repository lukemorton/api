package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()

	err := users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)
	assertUserStored(t, users)
}

func TestAutoIncrementID(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()

	users.Create(validUserWithDates("a@gmail.com"))
	users.Create(validUserWithDates("b@gmail.com"))
	assertIncrementedID(t, users)
}

func TestUniqueEmail(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()

	var err error

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.EqualError(t, err, "Email already taken")
}

func TestCreatePanicOnSQLError(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()
	users.createQuery = "hmm"

	assert.Panics(t, func() {
		users.Create(validUserWithDates("a@gmail.com"))
	})
}

func TestUpdateResetTokenHash(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()
	user := validUserWithDates("a@gmail.com")
	users.Create(user)

	user.ResetTokenHash = "bob"
	err := users.UpdateResetTokenHash(user)

	assert.Nil(t, err)
	assertResetPasswordTokenChanged(t, users)
}

func TestFindByEmail(t *testing.T) {
	users := SQLUserStore()
	users.CreateStore()
	users.Create(validUserWithDates("a@gmail.com"))

	user, err := users.FindByEmail("a@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "a@gmail.com", user.Email)
}

func TestFindByEmailWithUnknownEmail(t *testing.T) {
	users := SQLUserStore()
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

func assertUserStored(t *testing.T, db *sqlUserStore) {
	user := User{}
	db.Get(&user, "SELECT email, password_hash FROM users")
	assert.Equal(t, "a@gmail.com", user.Email)
	assert.Equal(t, "bob", user.PasswordHash)
}

func assertIncrementedID(t *testing.T, db *sqlUserStore) {
	var id []int
	err := db.Select(&id, "SELECT id FROM users")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(id))
	assert.Equal(t, 1, id[0])
	assert.Equal(t, 2, id[1])
}

func assertResetPasswordTokenChanged(t *testing.T, db *sqlUserStore) {
	var token string
	db.Get(&token, "SELECT reset_token_hash FROM users")
	assert.Equal(t, "bob", token)
}
