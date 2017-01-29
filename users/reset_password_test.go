package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetPassword(t *testing.T) {
	token, err := ResetPassword(mockUserPasswordResetter{}, ResetPasswordUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestResetPasswordWithMissingEmail(t *testing.T) {
	token, err := ResetPassword(mockUserPasswordResetter{}, ResetPasswordUser{})

	assert.Empty(t, token)
	assert.EqualError(t, err, "Email address required to reset password")
}

func TestResetPasswordWithInvalidEmail(t *testing.T) {
	_, err := ResetPassword(mockUserPasswordResetter{errors.New("Email not found")}, ResetPasswordUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.EqualError(t, err, "Email not found")
}

type mockUserPasswordResetter struct {
	err error
}

func (users mockUserPasswordResetter) UpdateResetTokenHashByEmail(email string, token string) error {
	return users.err
}
