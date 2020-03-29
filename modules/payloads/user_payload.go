package payloads

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type UserPayload struct {
	Name     string
	Email    string
	ImageUrl string
}

func (u *UserPayload) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&u.Email, validation.Required, validation.Length(1, 50)),
		validation.Field(&u.ImageUrl, validation.Length(0, 100)),
	)
}
