package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func LoadEmailConfig() *EmailConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &EmailConfig{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
