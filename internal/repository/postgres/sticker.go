package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"skin-monkey/internal/entity"
	"strings"
)

type StickerPostgres struct {
	db *sqlx.DB
}

func NewStickerPostgres(db *sqlx.DB) *StickerPostgres {
	return &StickerPostgres{
		db: db,
	}
}

func (r *StickerPostgres) CreateSticker(sticker *entity.Sticker) error {
	query := fmt.Sprintf(`INSERT INTO %s (instanceid, name, hash_name, sell_price, sell_price_text, is_custom_sell_price) VALUES ($1, $2, $3, $4, $5, false) returning instanceid`, stickersTable)

	_, err := r.db.Exec(query, sticker.InstanceId, sticker.Name, sticker.HashName, sticker.SellPrice, sticker.SellPriceText)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			err := r.UpdateSticker(sticker)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (r *StickerPostgres) GetStickerById(id string) (*entity.Sticker, error) {
	query := fmt.Sprintf(`SELECT instanceid, name, hash_name, sell_price, sell_price_text, is_custom_sell_price FROM %s WHERE instanceid = $1`, stickersTable)

	var sticker entity.Sticker

	err := r.db.Get(&sticker, query, id)
	if err != nil {
		return nil, err
	}

	return &sticker, nil
}

func (r *StickerPostgres) GetStickerByName(name string) (*entity.Sticker, error) {
	query := fmt.Sprintf(`SELECT instanceid, name, hash_name, sell_price, sell_price_text, is_custom_sell_price FROM %s WHERE hash_name = $1`, stickersTable)

	var sticker entity.Sticker

	err := r.db.Get(&sticker, query, name)
	if err != nil {
		return nil, err
	}

	return &sticker, nil
}

func (r *StickerPostgres) UpdateSticker(sticker *entity.Sticker) error {
	query := fmt.Sprintf(`UPDATE %s SET sell_price = $1, sell_price_text = $2 WHERE instanceid = $3`, stickersTable)

	_, err := r.db.Exec(query, sticker.SellPrice, sticker.SellPriceText, sticker.InstanceId)
	if err != nil {
		return err
	}

	return nil
}
