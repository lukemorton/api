package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := store()
	err := users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)
	assertUserStored(t, users)
}

func TestAutoIncrementID(t *testing.T) {
	users := store()
	users.Create(validUserWithDates("a@gmail.com"))
	users.Create(validUserWithDates("b@gmail.com"))
	assertIncrementedID(t, users)
}

func TestUniqueEmail(t *testing.T) {
	users := store()

	var err error

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.Nil(t, err)

	err = users.Create(validUserWithDates("a@gmail.com"))
	assert.EqualError(t, err, "Email already taken")
}

func TestCreatePanicOnSQLError(t *testing.T) {
	users := store()
	users.createQuery = "hmm"

	assert.Panics(t, func() {
		users.Create(validUserWithDates("a@gmail.com"))
	})
}

func TestUpdateResetTokenHash(t *testing.T) {
	users, user := storeAndUser()
	user.ResetTokenHash = "bob"
	users.UpdateResetTokenHash(&user)
	assertResetPasswordTokenChanged(t, users)
}

func TestUpdateResetTokenHashUpdatesUpdatedAt(t *testing.T) {
	users, user := storeAndUser()
	prevUpdatedAt := user.UpdatedAt
	users.UpdateResetTokenHash(&user)
	assert.NotEqual(t, prevUpdatedAt, user.UpdatedAt)
}

func TestUpdatePasswordHash(t *testing.T) {
	users, user := storeAndUser()
	user.PasswordHash = "bob"
	users.UpdatePasswordHash(&user)
	assertPasswordHashChanged(t, users)
}

func TestUpdatePasswordHashUpdatesUpdatedAt(t *testing.T) {
	users, user := storeAndUser()
	prevUpdatedAt := user.UpdatedAt
	users.UpdatePasswordHash(&user)
	assert.NotEqual(t, prevUpdatedAt, user.UpdatedAt)
}

func TestFindByEmail(t *testing.T) {
	users, _ := storeAndUser()
	user, err := users.FindByEmail("a@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "a@gmail.com", user.Email)
}

func TestFindByEmailWithUnknownEmail(t *testing.T) {
	users := store()
	_, err := users.FindByEmail("a@gmail.com")
	assert.EqualError(t, err, "Email not recognised")
}

func store() *sqlUserStore {
	users := SQLUserStore()
	users.CreateStore()
	users.ClearStore()
	return users
}

func storeAndUser() (*sqlUserStore, User) {
	users := store()
	user := validUserWithDates("a@gmail.com")
	users.Create(user)
	return users, *user
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
	db.Get(&user, "SELECT email, password_hash, created_at, updated_at FROM users")
	assert.Equal(t, "a@gmail.com", user.Email)
	assert.Equal(t, "bob", user.PasswordHash)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
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

func assertPasswordHashChanged(t *testing.T, db *sqlUserStore) {
	var token string
	db.Get(&token, "SELECT password_hash FROM users")
	assert.Equal(t, "bob", token)
}
