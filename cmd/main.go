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
	"skin-monkey/internal/lib/logger"
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

	bot := app.RunBot()

	repo := repository.NewRepository(db)
	services := service.NewService(repo, log, bot)
	handlers := handler.NewHandler(services, log)
	application := app.NewApp(log)

	go application.Run(handlers.InitRoutes(), cfg.App.Port, services.Bot)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.StopBot(services)
	application.Stop()
}

func initLogger() *slog.Logger {
	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := logger.NewPrettyHandler(os.Stdout, opts)
	log := slog.New(handler)

	return log
}
