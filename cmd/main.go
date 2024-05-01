package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	"os/signal"
	"skin-monkey/internal/app"
	"skin-monkey/internal/config"
	"skin-monkey/internal/handler"
	"skin-monkey/internal/lib/discordBot"
	"skin-monkey/internal/lib/logger"
	"skin-monkey/internal/lib/tgBot"
	repository "skin-monkey/internal/repository/postgres"
	"skin-monkey/internal/service"
	"syscall"
)

func main() {
	cfg := config.MustLoadConfig()

	log := initLogger()

	log.Info("Config loaded",
		slog.String("port", cfg.App.Port),
		slog.String("env", cfg.App.Env),
	)

	gin.SetMode(cfg.App.Env)

	db, err := repository.ConnectDb(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		DBname:   cfg.DB.DBname,
		SSL:      cfg.DB.SSL,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db: %s", err))
	}

	repo := repository.NewRepository(db)

	tgBot := InitTgBot(log, cfg.TgBot.Token, repo)
	discordBot := InitDiscordBot(log, cfg.DiscordBot.Token, cfg.DiscordBot.ChannelId)

	services := service.NewService(repo, log, tgBot, discordBot)
	handlers := handler.NewHandler(services, log)
	application := app.NewApp(log)

	go application.Run(handlers.InitRoutes(), cfg.App.Port, tgBot)
	go tgBot.Start()
	go discordBot.Start()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	err = application.Stop(tgBot, discordBot)
	if err != nil {
		return
	}
}

func InitTgBot(log *slog.Logger, botToken string, repo *repository.Repository) *tgBot.TgBotStruct {
	return tgBot.NewBot(log, botToken, repo)
}

func InitDiscordBot(log *slog.Logger, botToken string, ChannelId string) *discordBot.DiscordBotStruct {
	return discordBot.NewBot(log, botToken, ChannelId)
}

func initLogger() *slog.Logger {
	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	log := slog.New(logger.NewPrettyHandler(os.Stdout, opts))

	return log
}
