package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DSN         string
	HOSTNAME    string
	DB_HOST     string
	DB_NAME     string
	DB_PASSWORD string
	DB_USER     string
}

func LoadConfig() Config {

	godotenv.Load()

	if _, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Fatal("invalid PORT number")
	}

	return Config{
		Port:        os.Getenv("PORT"),
		DSN:         os.Getenv("DSN"),
		HOSTNAME:    os.Getenv("HOSTNAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_USER:     os.Getenv("DB_USER"),
	}
}
