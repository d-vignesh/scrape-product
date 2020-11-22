package utils

import (
	"github.com/hashicorp/go-hclog"
)

func NewLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name: 	"product-scraper",
		Level: 	hclog.LevelFromString("DEBUG"),
	})
}