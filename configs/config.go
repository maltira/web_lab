package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string

	AppPort string
	AppHost string

	Secret string
}

var Cfg *Config

type UserGroup string

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Не удалось получить доступ к .env файлу")
	}

	Cfg = &Config{
		DbHost:  os.Getenv("DB_HOST"),
		DbPort:  os.Getenv("DB_PORT"),
		DbUser:  os.Getenv("DB_USER"),
		DbPass:  os.Getenv("DB_PASS"),
		DbName:  os.Getenv("DB_NAME"),
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
		Secret:  os.Getenv("SECRET"),
	}
}
