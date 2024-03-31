package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"skin-monkey/internal/entity"
	"strings"
)

var (
	ErrSkinNotFound = errors.New("skin not found")
)

type SkinPostgres struct {
	db *sqlx.DB
}

func NewSkinPostgres(db *sqlx.DB) *SkinPostgres {
	return &SkinPostgres{
		db: db,
	}
}

func (r *SkinPostgres) CreateSkin(skin *entity.Skin) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, image, inspect_link, float, new, price, tradable, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`, skinsTable)

	_, err := r.db.Exec(query, skin.Id, skin.Name, skin.Image, skin.InspectLink, skin.Float, true, skin.Price, skin.Tradable, skin.Url)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			skin.New = false
			err := r.UpdateSkin(skin)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (r *SkinPostgres) GetSkin(id string) (*entity.Skin, error) {
	query := fmt.Sprintf(`SELECT id, name, image, inspect_link, float, new, price, tradable FROM %s WHERE id = $1`, skinsTable)

	var skin entity.Skin

	err := r.db.Get(&skin, query, id)
	if err != nil {
		return nil, err
	}

	return &skin, nil
}

func (r *SkinPostgres) UpdateSkin(skin *entity.Skin) error {
	query := fmt.Sprintf(`UPDATE %s SET new = $1 WHERE id = $2`, skinsTable)

	_, err := r.db.Exec(query, skin.New, skin.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SkinPostgres) DeleteSkin(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, skinsTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SkinPostgres) GetSkinsFilter() (*entity.Skin, error) {
	query := fmt.Sprintf(`SELECT id, name, image, inspect_link, float, new, price, tradable FROM %s`, skinsTable)

	var skin entity.Skin

	err := r.db.Get(&skin, query)
	if err != nil {
		return nil, err
	}

	return &skin, nil
}
