package utils

import (
	"github.com/hashicorp/go-hclog"
)

func NewLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name : "product-store",
		Level: hclog.LevelFromString("DEBUG"),
	})
}