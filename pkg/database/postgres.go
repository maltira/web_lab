package database

import (
	"fmt"
	config "web-lab/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Cfg.DbHost,
		config.Cfg.DbUser,
		config.Cfg.DbPass,
		config.Cfg.DbName,
		config.Cfg.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Не удалось подключиться к DB: %v", err))
	}

	return db
}
