package users

import (
	"errors"
)

type ResetPasswordUser struct {
	Email string `json:"email"`
}

func ResetPassword(users UserPasswordResetter, r ResetPasswordUser) (string, error) {
	err := validateResetPasswordUser(r)

	if err != nil {
		return "", err
	}

	user, err := users.FindByEmail(r.Email)

	if err != nil {
		return "", err
	}

	token := user.GenerateResetToken()
	return token, users.UpdateResetTokenHash(&user)
}

func validateResetPasswordUser(user ResetPasswordUser) error {
	if user.Email == "" {
		return errors.New("Email address required to reset password")
	} else {
		return nil
	}
}
