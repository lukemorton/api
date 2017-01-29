package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetPassword(t *testing.T) {
	token, err := ResetPassword(mockUserUpdater{}, ResetPasswordUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func TestResetPasswordWithMissingEmail(t *testing.T) {
	token, err := ResetPassword(mockUserUpdater{}, ResetPasswordUser{})

	assert.Empty(t, token)
	assert.EqualError(t, err, "Email address required to reset password")
}

func TestResetPasswordWithInvalidEmail(t *testing.T) {
	_, err := ResetPassword(mockUserUpdater{errors.New("Email not found")}, ResetPasswordUser{
		Email: "lukemorton.dev@gmail.com",
	})

	assert.EqualError(t, err, "Email not found")
}

type mockUserUpdater struct {
	err error
}

func (users mockUserUpdater) UpdateResetTokenHashByEmail(email string, token string) error {
	return users.err
}
