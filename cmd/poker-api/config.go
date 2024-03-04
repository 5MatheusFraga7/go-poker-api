package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	APIKey   string
	Database string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	apiKey := os.Getenv("API_KEY")
	database := os.Getenv("DATABASE")

	return Config{
		Host:     host,
		Port:     port,
		APIKey:   apiKey,
		Database: database,
	}
}
