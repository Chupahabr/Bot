package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	App        AppConfig
	DB         DBConfig
	TgBot      BotConfig
	DiscordBot BotConfig
}

type AppConfig struct {
	Port string
	Env  string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSL      string
}

type BotConfig struct {
	Token     string
	ChannelId string
}

func MustLoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file: " + err.Error())
	}

	var cfg Config

	cfg.App.Port = getEnv("APP_PORT")
	cfg.App.Env = getEnv("APP_ENV")
	cfg.DB.Host = getEnv("DB_HOST")
	cfg.DB.Port = getEnv("DB_PORT")
	cfg.DB.Username = getEnv("DB_USER")
	cfg.DB.Password = getEnv("DB_PASSWORD")
	cfg.DB.DBname = getEnv("DB_NAME")
	cfg.DB.SSL = getEnv("DB_SSL")
	cfg.TgBot.Token = getEnv("TG_BOT_TOKEN")
	cfg.DiscordBot.Token = getEnv("DISCORD_BOT_TOKEN")
	cfg.DiscordBot.ChannelId = getEnv("SEND_INSPECT_URL_CHANEL")

	if err = cleanenv.ReadEnv(&cfg); err != nil {
		panic("failed to load config from environment: " + err.Error())
	}

	return &cfg
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		panic("environment variable not found: " + key)
	}

	return value
}
