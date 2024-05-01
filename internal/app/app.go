package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"skin-monkey/internal/lib/discordBot"
	"skin-monkey/internal/lib/tgBot"
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

func (a *App) Run(handler http.Handler, port string, bot *tgBot.TgBotStruct) error {
	a.log.Info("Starting server")

	a.HttpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	err := bot.SendText("Бот запущен")
	if err != nil {
		return err
	}

	err = a.HttpServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%s: %w", "Failed to start server", err)
	}

	return nil
}

func (a *App) Stop(bot *tgBot.TgBotStruct, discordBot *discordBot.DiscordBotStruct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		a.log.Error("Graceful shutdown failed:", err)
		return err
	}
	err := a.StopBot(bot)
	if err != nil {
		return err
	}

	discordBot.Stop()
	a.log.Info("Server gracefully stopped")

	return nil
}

func (a *App) StopBot(bot *tgBot.TgBotStruct) error {
	err := bot.SendText("Бот остановлен")
	if err != nil {
		return err
	}
	return nil
}
