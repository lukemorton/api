package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetPassword(t *testing.T) {
	token, err := ResetPassword(mockPasswordResetterFindUser(), ResetPasswordRequest{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestResetPasswordWithMissingEmail(t *testing.T) {
	token, err := ResetPassword(mockPasswordResetterFindUser(), ResetPasswordRequest{})

	assert.Empty(t, token)
	assert.EqualError(t, err, "Email address required to reset password")
}

func TestResetPasswordWithInvalidEmail(t *testing.T) {
	_, err := ResetPassword(mockPasswordResetterFindErr("Email not found"), ResetPasswordRequest{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.EqualError(t, err, "Email not found")
}

func mockPasswordResetterFindUser() mockUserPasswordResetter {
	return mockUserPasswordResetter{
		findUser: User{},
	}
}

func mockPasswordResetterFindErr(err string) mockUserPasswordResetter {
	return mockUserPasswordResetter{
		findErr: errors.New(err),
	}
}

type mockUserPasswordResetter struct {
	findUser User
	findErr  error
}

func (users mockUserPasswordResetter) FindByEmail(email string) (User, error) {
	return users.findUser, users.findErr
}

func (users mockUserPasswordResetter) UpdateResetTokenHash(user *User) {
}
