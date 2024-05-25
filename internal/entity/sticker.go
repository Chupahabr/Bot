package entity

import "github.com/go-playground/validator/v10"

type Sticker struct {
	InstanceId        string `json:"id" db:"instanceid"`
	Name              string `json:"name" db:"name"`
	HashName          string `json:"hashName" db:"hash_name"`
	SellPrice         int    `json:"sellPrice" db:"sell_price"`
	SellPriceText     string `json:"sellPriceText" db:"sell_price_text"`
	IsCustomSellPrice bool   `json:"isCustomSellPrice" db:"is_custom_sell_price"`
}

func (u *Sticker) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
