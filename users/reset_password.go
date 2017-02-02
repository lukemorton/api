package users

import (
	"errors"
)

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func ResetPassword(users UserPasswordResetter, r ResetPasswordRequest) (string, error) {
	err := validateResetPasswordRequest(r)

	if err != nil {
		return "", err
	}

	user, err := users.FindByEmail(r.Email)

	if err != nil {
		return "", err
	}

	token := user.GenerateResetToken()
	users.UpdateResetTokenHash(&user)
	return token, nil
}

func validateResetPasswordRequest(r ResetPasswordRequest) error {
	if r.Email == "" {
		return errors.New("Email address required to reset password")
	} else {
		return nil
	}
}
