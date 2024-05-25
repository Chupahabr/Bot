package cron

import (
	"log/slog"
	repository "skin-monkey/internal/repository/postgres"
	"time"
)

type Cron struct {
	log        *slog.Logger
	repo       *repository.Repository
	cronEvents map[string]int64
}

func NewCron(log *slog.Logger, repo *repository.Repository) *Cron {
	return &Cron{
		log:  log,
		repo: repo,
		cronEvents: map[string]int64{
			"ParseStickers": 0,
		},
	}
}

func (c *Cron) Start() {
	for {
		for cronName, lastTrigger := range c.cronEvents {
			if lastTrigger <= time.Now().Unix() {
				if cronName == "ParseStickers" {
					go c.ParseStickers()
				}
				c.cronEvents[cronName] = time.Now().Add(24 * time.Hour).Unix()
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

func (c *Cron) ParseStickers() {
	c.log.Info("ParseStickers")

	/*stickerParser := parser.NewStickerParser(c.log, c.repo)
	stickerParser.Run()*/
}
