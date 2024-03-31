package service

import (
	"log/slog"
	"skin-monkey/internal/entity"
	"skin-monkey/internal/lib/bot"
	"skin-monkey/internal/repository/postgres"
)

type Skin interface {
	CreateSkin(*entity.Skin) error
	UpdateSkin(*entity.Skin) error
	GetSkin(string) (*entity.Skin, error)
	DeleteSkin(string) error
	GetSkinFilter() (*entity.Skin, error)
}

type Service struct {
	Skin
}

func NewService(repo *repository.Repository, log *slog.Logger, botStruct *bot.BotStruct) *Service {
	return &Service{
		Skin: NewSkinService(repo.Skin, log, botStruct),
	}
}
