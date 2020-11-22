package utils

import (
	"os"
)

type Configuration struct {
	ServerAddress string
}

func NewConfiguration() *Configuration {

	config := Configuration {
		ServerAddress: 	os.Getenv("SCRAPER_SERVER_ADDRESS"),
	}

	return &config
}