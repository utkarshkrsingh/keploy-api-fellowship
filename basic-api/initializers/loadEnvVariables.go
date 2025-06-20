package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	if godotenv.Load() != nil {
		log.Fatal("Failed to load environment variables")
	}
}
