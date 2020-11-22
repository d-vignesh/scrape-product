package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-scraper/utils"
	"github.com/d-vignesh/scrape-product/product-scraper/services"
)

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

type ScrapeRequest struct {
	url string  `json:"url" validate:"required"`
}

type GenericError struct {
	Message string	`json:"message"`
}