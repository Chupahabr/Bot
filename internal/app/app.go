package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"net/http"
	"skin-monkey/internal/service"
	"time"
)

type App struct {
	HttpServer *http.Server
	log        *slog.Logger
}

func NewApp(log *slog.Logger) *App {
	return &App{
		HttpServer: nil,
		log:        log,
	}
}

func (a *App) Run(handler http.Handler, port string, botService *service.BotService) error {
	a.log.Info("Starting server")

	a.HttpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	botService.SendText("Бот запущен")

	err := a.HttpServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%s: %w", "Failed to start server", err)
	}

	return nil
}

func (a *App) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		a.log.Error("Graceful shutdown failed:", err)
		return err
	}
	a.log.Info("Server gracefully stopped")

	return nil
}

func RunBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("5931349262:AAHQGV4ivSuKsu8HvMEN05-v5qK7siduF4E")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func (a *App) StopBot(services *service.Service) error {
	services.Bot.SendText("Бот остановлен")
	return nil
}
