package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePasswordByPassword(t *testing.T) {
	user, err := ChangePassword(mockPasswordChangerFindUser(), ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		Password: "bob",
		NewPassword: "fred",
	})

	assert.Nil(t, user.VerifyPassword("fred"))
	assert.Nil(t, err)
}

func TestChangePasswordByPasswordWithMissingNewPassword(t *testing.T) {
	_, err := ChangePassword(mockUserPasswordChanger{}, ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		Password: "bob",
	})

	assert.EqualError(t, err, "New password required to change password")
}

func TestChangePasswordByResetToken(t *testing.T) {
	user, token := userWithResetToken()

	user, err := ChangePassword(mockUserPasswordChanger{findUser: user}, ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		ResetToken: token,
		NewPassword: "fred",
	})

	assert.Nil(t, user.VerifyPassword("fred"))
	assert.Nil(t, err)
}

func TestChangePasswordByResetTokenWithMissingNewPassword(t *testing.T) {
	user, token := userWithResetToken()

	user, err := ChangePassword(mockUserPasswordChanger{findUser: user}, ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		ResetToken: token,
	})

	assert.EqualError(t, err, "New password required to change password")
}

func TestChangePasswordWithMissingPasswordAndResetToken(t *testing.T) {
	_, err := ChangePassword(mockUserPasswordChanger{}, ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		NewPassword: "bob",
	})

	assert.EqualError(t, err, "Password or reset token required to change password")
}

func TestChangePasswordWithMissingEmail(t *testing.T) {
	_, err := ChangePassword(mockPasswordChangerFindUser(), ChangePasswordRequest{
		Password: "bob",
		NewPassword: "fred",
	})

	assert.EqualError(t, err, "Email address required to change password")
}

func TestChangePasswordWithInvalidEmail(t *testing.T) {
	_, err := ChangePassword(mockPasswordChangerFindErr("Email not found"), ChangePasswordRequest{
		Email: "lukemorton.dev@gmail.com",
		Password: "bob",
		NewPassword: "fred",
	})

	assert.EqualError(t, err, "Email not found")
}

func userWithResetToken() (User, string) {
	user := User{}
	return user, user.GenerateResetToken()
}

func mockPasswordChangerFindUser() mockUserPasswordChanger {
	user := User{}
	user.SetPassword("bob")

	return mockUserPasswordChanger{
		findUser: user,
	}
}

func mockPasswordChangerFindErr(err string) mockUserPasswordChanger {
	return mockUserPasswordChanger{
		findErr: errors.New(err),
	}
}

type mockUserPasswordChanger struct {
	findUser User
	findErr  error
}

func (users mockUserPasswordChanger) FindByEmail(email string) (User, error) {
	return users.findUser, users.findErr
}

func (users mockUserPasswordChanger) UpdatePasswordHash(user *User) {
}
