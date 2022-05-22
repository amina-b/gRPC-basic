package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config Configuration

func Load() error {

	err := godotenv.Load()

	if err != nil {
		log.Printf("Error occured while loading .env file. Err: %s", err)
		return err
	}

	Config = Configuration{
		Environment: Environment{
			Port:    os.Getenv("PORT"),
			Address: os.Getenv("ADDRESS"),
		},
	}

	return nil

}
