package utils

import (
	"os"
)

type Configuration struct {
	// address to start the scraper service
	ServerAddress string
	// address of the product store service
	StoreServerAddress string
}

func NewConfiguration() *Configuration {

	config := Configuration {
		ServerAddress: 	os.Getenv("SCRAPER_SERVER_ADDRESS"),
		StoreServerAddress: os.Getenv("STORE_SERVER_ADDRESS"),
	}

	return &config
}