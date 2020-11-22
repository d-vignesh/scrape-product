package utils

import (
	"os"
)

type Configuration struct {
	// address to start the Store service
	ServerAddress string
	// database details
	DBHost		  string
	DBName		  string
	DBUser		  string
	DBPass		  string
	DBPort		  string
}

func NewConfiguration() *Configuration {

	configs := &Configuration {
		ServerAddress : os.Getenv("STORE_SERVER_ADDRESS"),
		DBHost 	: os.Getenv("DB_HOST"),
		DBName  : os.Getenv("DB_NAME"),
		DBUser  : os.Getenv("DB_USER"),
		DBPass  : os.Getenv("DB_PASS"),
		DBPort	: os.Getenv("DB_PORT"),
	}

	return configs
}