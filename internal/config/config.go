package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DSN      string
	HOSTNAME string
}

func LoadConfig() Config {

	godotenv.Load()

	if _, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Fatal("invalid PORT number")
	}

	return Config{
		Port:     os.Getenv("PORT"),
		DSN:      os.Getenv("DSN"),
		HOSTNAME: os.Getenv("HOSTNAME"),
	}
}
