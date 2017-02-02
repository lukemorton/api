package users

import (
	"errors"
)

type ChangePasswordUser struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	ResetToken  string `json:"reset_token"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(users UserPasswordUpdater, c ChangePasswordUser) (User, error) {
	err := validateChangePasswordUser(c)

	if err != nil {
		return User{}, err
	}

	if c.Password != "" {
		return changePasswordByEmailAndPassword(users, c)
	} else if c.ResetToken != "" {
		return changePasswordByEmailAndResetToken(users, c)
	} else {
		panic("Should not be here")
	}
}

func validateChangePasswordUser(user ChangePasswordUser) error {
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

func changePasswordByEmailAndPassword(users UserPasswordUpdater, c ChangePasswordUser) (user User, err error) {
	user, err = users.FindByEmail(c.Email)

	if err != nil {
		return
	}

	err = user.VerifyPassword(c.Password)

	if err != nil {
		return
	}

	return setPassword(users, user, c.NewPassword), nil
}

func changePasswordByEmailAndResetToken(users UserPasswordUpdater, c ChangePasswordUser) (user User, err error) {
	user, err = users.FindByEmail(c.Email)

	if err != nil {
		return
	}

	err = user.VerifyResetToken(c.ResetToken)

	if err != nil {
		return
	}

	return setPassword(users, user, c.NewPassword), nil
}

func setPassword(users UserPasswordUpdater, user User, newPassword string) User {
	user.SetPassword(newPassword)
	users.UpdatePasswordHash(&user)
	return user
}
