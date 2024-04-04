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

type Repository struct {
	Skin
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Skin: NewSkinPostgres(db),
		User: NewUserPostgres(db),
	}
}
