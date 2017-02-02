package users

import (
	"errors"
)

type ChangePasswordRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	ResetToken  string `json:"reset_token"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(users UserPasswordUpdater, r ChangePasswordRequest) (User, error) {
	err := validateChangePasswordRequest(r)

	if err != nil {
		return User{}, err
	}

	if r.Password != "" {
		return changePasswordByEmailAndPassword(users, r)
	} else if r.ResetToken != "" {
		return changePasswordByEmailAndResetToken(users, r)
	} else {
		panic("Should not be here")
	}
}

func validateChangePasswordRequest(user ChangePasswordRequest) error {
	if user.Email == "" {
		return errors.New("Email address required to change password")
	} else if user.Password == "" && user.ResetToken == "" {
		return errors.New("Password or reset token required to change password")
	} else if user.NewPassword == "" {
		return errors.New("New password required to change password")
	} else {
		return nil
	}
}

func changePasswordByEmailAndPassword(users UserPasswordUpdater, r ChangePasswordRequest) (user User, err error) {
	user, err = users.FindByEmail(r.Email)

	if err != nil {
		return
	}

	err = user.VerifyPassword(r.Password)

	if err != nil {
		return
	}

	return setPassword(users, user, r.NewPassword), nil
}

func changePasswordByEmailAndResetToken(users UserPasswordUpdater, r ChangePasswordRequest) (user User, err error) {
	user, err = users.FindByEmail(r.Email)

	if err != nil {
		return
	}

	err = user.VerifyResetToken(r.ResetToken)

	if err != nil {
		return
	}

	return setPassword(users, user, r.NewPassword), nil
}

func setPassword(users UserPasswordUpdater, user User, newPassword string) User {
	user.SetPassword(newPassword)
	users.UpdatePasswordHash(&user)
	return user
}
