package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"skin-monkey/internal/entity"
	"skin-monkey/internal/repository/postgres"
)

type Skin interface {
	CreateSkin(*entity.Skin) error
	UpdateSkin(*entity.Skin) error
	GetSkin(string) (*entity.Skin, error)
	DeleteSkin(string) error
	GetSkinFilter() (*entity.Skin, error)
}

type Bot interface {
	SendText(text string) error
}

type Service struct {
	Skin
	Bot
}

func NewService(repo *repository.Repository, log *slog.Logger, bot *tgbotapi.BotAPI) *Service {
	BotService := NewBotService(bot, log)
	return &Service{
		Skin: NewSkinService(repo.Skin, log, BotService),
		Bot:  *BotService,
	}
}
