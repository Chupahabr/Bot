package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type BotStruct struct {
	log *slog.Logger
	Bot *tgbotapi.BotAPI
}

func NewBot(log *slog.Logger) *BotStruct {
	botObject, err := tgbotapi.NewBotAPI("5931349262:AAHQGV4ivSuKsu8HvMEN05-v5qK7siduF4E")
	if err != nil {
		panic(err)
	}
	botObject.Debug = true

	fmt.Printf("Authorized on account %s", botObject.Self.UserName)

	return &BotStruct{
		log,
		botObject,
	}
}

func (b BotStruct) SendText(text string) error {
	// ruslan 693559920
	// sanya 1064622908

	//chatId1 := 693559920
	//chatId2 := 1064622908

	msg := tgbotapi.NewMessage(693559920, text)
	_, err := b.Bot.Send(msg)
	if err != nil {
		return err
	}

	msg = tgbotapi.NewMessage(1064622908, text)
	_, err = b.Bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
