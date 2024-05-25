package entity

import (
	"github.com/go-playground/validator/v10"
)

type Skin struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"marketHashName" db:"name"`
	Image       string `json:"image" db:"image"`
	InspectLink string `json:"inspectLink" db:"inspect_link"`
	Float       string `json:"float" db:"float"`
	New         bool   `json:"new" db:"new"`
	Price       string `json:"sellPrice" db:"price"`
	Tradable    bool   `json:"tradable" db:"tradable"`
	Url         string `json:"url" db:"url"`
	Stickers    []StickerShort
}

type StickerShort struct {
	Name  string `json:"name" db:"name"`
	Image string `json:"image" db:"image"`
}

func (u *Skin) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
