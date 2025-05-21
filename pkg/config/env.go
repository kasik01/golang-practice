package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	JWT_SECRET  string
	APP_PORT    string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

var appConfig *AppConfig

func LoadEnv() {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	// khi chạy dev environment
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	appConfig = &AppConfig{
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		APP_PORT:    os.Getenv("APP_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
	}

	fmt.Println("Environment variables loaded successfully")
}

func GetAppConfig() *AppConfig {
	return appConfig
}
