package initializers

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	envPath := filepath.Join(".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Failed to load env variables: %v", err)
	}
}
