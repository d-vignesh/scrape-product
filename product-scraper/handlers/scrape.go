package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-scraper/utils"
	"github.com/d-vignesh/scrape-product/product-scraper/services"
)

// ScrapingHandler contains instance of logger, config, scraper which are required for scraping service
type ScrapingHandler struct {
	logger 	hclog.Logger 
	config 	*utils.Configuration
	scraper *services.Scraper 
}

func NewScrapingHandler(logger hclog.Logger, config *utils.Configuration, scraper *services.Scraper) *ScrapingHandler {
	return &ScrapingHandler{
		logger : logger,
		config : config,
		scraper: scraper,
	}
}

// GenericError is a generic error message returned by server
type GenericError struct {
	Message string	`json:"message"`
}