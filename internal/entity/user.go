package entity

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id           int64  `json:"id" db:"id"`
	UserName     string `json:"user_name" db:"user_name"`
	Name         string `json:"name" db:"name"`
	LanguageCode string `json:"language_code" db:"language_code"`
	IsPremium    bool   `json:"is_premium" db:"is_premium"`
	IsBot        bool   `json:"is_bot" db:"is_bot"`
	Active       bool   `json:"active" db:"active"`
	DateAdd      int    `json:"date_add" db:"date_add"`
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
