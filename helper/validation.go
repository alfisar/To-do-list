package helper

import (
	"todolist/domain"

	validator "todolist/internal/validation"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidationDataUser(data domain.User) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.Name, validator.Required, validator.AlphanumericSimbols),
		validation.Field(&data.Username, validator.Required, validator.AlphanumericSimbols),
	)
	return
}

func ValidationLogin(data domain.Login) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
	)
	return
}
