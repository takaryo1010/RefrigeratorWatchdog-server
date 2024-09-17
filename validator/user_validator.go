package validator

import (
	"RefrigeratorWatchdog-server/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	ValidateUser(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) ValidateUser(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&user.Email, validation.Required, validation.Length(1, 255), is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(1, 255)),
	)
}
