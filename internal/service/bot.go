package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type BotService struct {
	log *slog.Logger
	bot *tgbotapi.BotAPI
}

func NewBotService(bot *tgbotapi.BotAPI, log *slog.Logger) *BotService {
	return &BotService{
		log,
		bot,
	}
}

func (b BotService) SendText(text string) error {
	// ruslan 693559920
	// sanya 1064622908

	//chatId1 := 693559920
	//chatId2 := 1064622908

	msg := tgbotapi.NewMessage(693559920, text)
	b.bot.Send(msg)

	/*msg = tgbotapi.NewMessage(int64(chatId2), text)
	bot.Send(msg)*/

	return nil
}
