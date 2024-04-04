package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"skin-monkey/internal/entity"
	repository "skin-monkey/internal/repository/postgres"
)

type BotStruct struct {
	log  *slog.Logger
	Bot  *tgbotapi.BotAPI
	repo *repository.Repository
}

func NewBot(log *slog.Logger, token string, repo *repository.Repository) *BotStruct {
	botObject, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	botObject.Debug = true

	fmt.Printf("Authorized on account %s", botObject.Self.UserName)

	return &BotStruct{
		log,
		botObject,
		repo,
	}
}

func (b BotStruct) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			user := entity.User{
				Id:           update.Message.From.ID,
				UserName:     update.Message.From.FirstName,
				Name:         update.Message.From.UserName,
				LanguageCode: update.Message.From.LanguageCode,
				IsBot:        update.Message.From.IsBot,
				DateAdd:      update.Message.Date,
				Active:       false,
			}

			b.repo.User.CreateUser(&user)

			fmt.Printf("[%s] %s \n chantId: %d \n", update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)
		}
	}
}

func (b BotStruct) SendText(text string) error {
	var users *[]entity.User

	users, _ = b.repo.User.GetUsersFilter()

	for _, user := range *users {
		msg := tgbotapi.NewMessage(user.Id, text)

		msg.ParseMode = "HTML"

		_, err := b.Bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
