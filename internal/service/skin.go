package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"skin-monkey/internal/entity"
	"skin-monkey/internal/lib/discordBot"
	"skin-monkey/internal/lib/tgBot"
	repository "skin-monkey/internal/repository/postgres"
	"strings"
	"time"
)

const (
	domainImages = "https://s.swap.gg/"
)

type SkinService struct {
	repository repository.Skin
	log        *slog.Logger
	tgBot      *tgBot.TgBotStruct
	discordBot *discordBot.DiscordBotStruct
}

func NewSkinService(repo repository.Skin, log *slog.Logger, TgBotStruct *tgBot.TgBotStruct, DiscordBotStruct *discordBot.DiscordBotStruct) *SkinService {
	return &SkinService{
		repository: repo,
		log:        log,
		tgBot:      TgBotStruct,
		discordBot: DiscordBotStruct,
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
		var image string

		if skin.InspectLink != "" {
			message, _ := s.discordBot.Bot.ChannelMessageSend(s.discordBot.ChannelId, skin.InspectLink)

			time.Sleep(10 * time.Second)

			messages, _ := s.discordBot.Bot.ChannelMessages(s.discordBot.ChannelId, 100, "", message.ID, message.ID)

			if len(messages) != 0 {
				for _, messageN := range messages {
					if messageN.ReferencedMessage != nil && messageN.ReferencedMessage.ID == message.ID {
						image = messageN.Content
					}
				}
			}
		}

		inspectLink := strings.ReplaceAll(skin.InspectLink, "%20", " ")

		var screenshotDataS tgBot.ResponseData
		screenshotData, _ := s.tgBot.ScreenshotRequest(inspectLink)
		screenshotDataS = screenshotData

		messageText := "Новый скин \n\n"

		messageText += fmt.Sprintf("Название: %s \n", skin.Name)
		messageText += fmt.Sprintf("Цена: <b>%s</b> руб \n\n", skin.Price)

		if image != "" {
			messageText += fmt.Sprintf("<a href='%s'>Изображение</a>\n", image)
		} else {
			messageText += fmt.Sprintf("<a href='%s'>Изображение</a> (Дефолтное)\n", skin.Image)
		}

		var tredable string
		if skin.Tradable {
			tredable = "Да"
		} else {
			tredable = "Нет"
		}

		messageText += fmt.Sprintf("<a href='%s'>Ссылка на скин</a> \n", skin.Url)
		messageText += fmt.Sprintf("Можно выкупить: %s \n", tredable)
		messageText += fmt.Sprintf("Флоат: %s \n\n", skin.Float)
		//messageText += fmt.Sprintf("[Стикер](%s) \n", "https://steamcdn-a.akamaihd.net/apps/730/icons/econ/stickers/stockh2021/liq_holo.b3bc7d3028b8e7214ee07c1b143b3e62522fbe54.png")

		if screenshotDataS.Status == "OK" {
			for _, value := range screenshotDataS.Result.Meta.Images {
				messageText += fmt.Sprintf("%s (Стето на %d%s. Позиция: %d) \n", value.Name, value.Wear, "%", value.Slot)
			}
		} else {
			for _, value := range skin.Stickers {
				messageText += fmt.Sprintf("<a href='%s'>%s</a> \n", value.Image, value.Name)
			}
		}

		messageText += fmt.Sprintf("\nСсылка на просмотр: %s \n", skin.InspectLink)

		err := s.tgBot.SendText(messageText)
		if err != nil {
			return err
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
