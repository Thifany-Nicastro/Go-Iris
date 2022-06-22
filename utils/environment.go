package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

func GetEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", name)
		return ""
	}

	return value
}
