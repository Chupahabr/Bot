package service

import (
	"database/sql"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"skin-monkey/internal/entity"
	"skin-monkey/internal/lib/bot"
	repository "skin-monkey/internal/repository/postgres"
)

type SkinService struct {
	repository repository.Skin
	log        *slog.Logger
	bot        *bot.BotStruct
}

func NewSkinService(repo repository.Skin, log *slog.Logger, botStruct *bot.BotStruct) *SkinService {
	return &SkinService{
		repository: repo,
		log:        log,
		bot:        botStruct,
	}
}

func (s *SkinService) CreateSkin(skin *entity.Skin) error {
	if err := skin.Validate(); err != nil {
		return err
	}

	err := s.repository.CreateSkin(skin)
	if err != nil {
		return err
	}

	skinDb, _ := s.repository.GetSkin(skin.Id)
	if skinDb.New {
		var tredable string
		if skin.Tradable {
			tredable = "Да"
		} else {
			tredable = "Нет"
		}

		messageText := "Новый скин \n\n"

		messageText += fmt.Sprintf("Название: %s \n", skin.Name)
		messageText += fmt.Sprintf("<a href='%s'>Ссылка на скин</a> \n", skin.Url)
		messageText += fmt.Sprintf("<a href='%s'>Изображение</a> \n", skin.Image)
		//messageText += fmt.Sprintf("Ссылка на просмотр: %s \n", skin.InspectLink)
		messageText += fmt.Sprintf("Можно выкупить: %s \n", tredable)
		messageText += fmt.Sprintf("Цена: %s руб \n\n", skin.Price)
		//messageText += fmt.Sprintf("[Стикер](%s) \n", "https://steamcdn-a.akamaihd.net/apps/730/icons/econ/stickers/stockh2021/liq_holo.b3bc7d3028b8e7214ee07c1b143b3e62522fbe54.png")

		for _, value := range skin.Stickers {
			messageText += fmt.Sprintf("<a href='%s'>%s</a> \n", value.Image, value.Name)
		}

		message := tgbotapi.NewMessage(693559920, messageText)
		message2 := tgbotapi.NewMessage(1064622908, messageText)
		message3 := tgbotapi.NewMessage(850418238, messageText)

		message.ParseMode = "HTML"
		message2.ParseMode = "HTML"
		message3.ParseMode = "HTML"

		_, err = s.bot.Bot.Send(message)
		if err != nil {
			log.Fatal(err)
		}

		_, err = s.bot.Bot.Send(message2)
		if err != nil {
			log.Fatal(err)
		}

		_, err = s.bot.Bot.Send(message3)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (s *SkinService) GetSkin(id string) (*entity.Skin, error) {
	return s.repository.GetSkin(id)
}

func (s *SkinService) UpdateSkin(skin *entity.Skin) error {
	existingSkin, err := s.GetSkin(skin.Id)
	if err != nil {
		return err
	}

	existingSkin.Name = skin.Name

	if err = existingSkin.Validate(); err != nil {
		s.log.Error(err.Error())
		return err
	}

	if err = s.repository.UpdateSkin(existingSkin); err != nil {
		return err
	}

	return nil
}

func (s *SkinService) DeleteSkin(id string) error {
	_, err := s.GetSkin(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrSkinNotFound
		}

		return err
	}

	return s.repository.DeleteSkin(id)
}

func (s *SkinService) GetSkinFilter() (*entity.Skin, error) {
	return s.repository.GetSkinsFilter()
}
