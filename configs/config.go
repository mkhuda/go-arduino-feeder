package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiHost string `env:"API_HOST"`
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading ENV")
	}

	return Config{
		ApiHost: os.Getenv("API_HOST"),
	}
}
