package service

import (
	"log/slog"
	"skin-monkey/internal/entity"
	"skin-monkey/internal/lib/discordBot"
	"skin-monkey/internal/lib/tgBot"
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

func NewService(repo *repository.Repository, log *slog.Logger, TgBotStruct *tgBot.TgBotStruct, DiscordBotStruct *discordBot.DiscordBotStruct) *Service {
	return &Service{
		Skin: NewSkinService(repo.Skin, repo.Sticker, log, TgBotStruct, DiscordBotStruct),
	}
}
