package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	DB_USER      = ""
	DB_PORT      = ""
	DB_PASSWORD  = ""
	DB_NAME      = ""
	DB_HOST      = ""
	DB_SSL       = "disable"
	DB_TIME_ZONE = ""
	PORT         = ""
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	PORT = os.Getenv("PORT")
	if os.Getenv("DB_SSL") != "" {
		DB_SSL = os.Getenv("DB_SSL")
	}

	DB_TIME_ZONE = os.Getenv("DB_TIMEZONE")
	return nil
}
