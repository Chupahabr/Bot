package repository

import (
	"github.com/jmoiron/sqlx"
	"skin-monkey/internal/entity"
)

type Skin interface {
	CreateSkin(user *entity.Skin) error
	GetSkin(string) (*entity.Skin, error)
	UpdateSkin(*entity.Skin) error
	DeleteSkin(string) error
	GetSkinsFilter() (*entity.Skin, error)
}

type User interface {
	CreateUser(user *entity.User) error
	GetUser(string) (*entity.User, error)
	GetUsersFilter() (*[]entity.User, error)
}

type Sticker interface {
	CreateSticker(sticker *entity.Sticker) error
	GetStickerById(id string) (*entity.Sticker, error)
	GetStickerByName(name string) (*entity.Sticker, error)
	UpdateSticker(sticker *entity.Sticker) error
}

type Repository struct {
	Skin
	User
	Sticker
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Skin:    NewSkinPostgres(db),
		User:    NewUserPostgres(db),
		Sticker: NewStickerPostgres(db),
	}
}
