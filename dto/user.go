package dto

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserDTO struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserDTO) ValidateUser() error {
	err := v.ValidateStruct(&u,
		v.Field(&u.Name, v.Required.Error("name is required")),
		v.Field(&u.Email, v.Required, is.Email.Error("invalid email")),
		v.Field(&u.Password, v.Required.Error("password is required"), v.Length(6, 12).Error("password must be between 6 and 12 characters")),
		v.Field(&u.ConfirmPassword, v.Required.Error("confirm password is required")),
	)

	if err != nil {
		return err
	}

	if u.Password != u.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}

	return nil
}

func (u UserLoginDTO) ValidateUserLogin() error {
	err := v.ValidateStruct(&u,
		v.Field(&u.Email, v.Required, is.Email.Error("invalid email")),
		v.Field(&u.Password, v.Required.Error("password is required")),
	)

	if err != nil {
		return err
	}

	return nil
}
