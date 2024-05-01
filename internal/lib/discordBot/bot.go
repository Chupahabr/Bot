package discordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

type DiscordBotStruct struct {
	log       *slog.Logger
	Bot       *discordgo.Session
	ChannelId string
}

func NewBot(log *slog.Logger, token string, ChannelId string) *DiscordBotStruct {
	// Создаем новую сессию DiscordGo
	discordBot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}

	return &DiscordBotStruct{
		log:       log,
		Bot:       discordBot,
		ChannelId: ChannelId,
	}
}

func (b DiscordBotStruct) Start() {
	// Открываем сессию соединения
	err := b.Bot.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
}

func (b DiscordBotStruct) Stop() {
	// Закрываем сессию соединения при выходе
	b.Bot.Close()
}
