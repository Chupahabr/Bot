package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"net/http"
	"skin-monkey/internal/entity"
	repository "skin-monkey/internal/repository/postgres"
)

const (
	domainSkreenshot = "https://api.swap.gg/"
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

type ResponseData struct {
	Status string `json:"status"`
	Result struct {
		ImageID     string `json:"imageId"`
		MarketName  string `json:"marketName"`
		InspectLink string `json:"inspectLink"`
		State       string `json:"state"`
		Meta        struct {
			Images []struct {
				Slot int    `json:"slot"`
				Name string `json:"name"`
				Wear int    `json:"wear"`
			} `json:"5"`
		} `json:"meta"`
	} `json:"result"`
}

func (b BotStruct) ScreenshotRequest(inspectLink string) (ResponseData, error) {
	url := domainSkreenshot + "v2/screenshot"

	data := map[string]string{
		"inspectLink": inspectLink,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ResponseData{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return ResponseData{}, err
	}
	defer resp.Body.Close()

	var responseData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return ResponseData{}, err
	}

	return responseData, nil
}
