package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-store/utils"
	"github.com/d-vignesh/scrape-product/product-store/data"
)

// ProductHandler contains instance of logger, config, repo and validator which are required for scraping service
type ProductHandler struct {
	logger 		hclog.Logger
	config 		*utils.Configuration
	repo   		data.Repository
	validator 	*data.Validation
}

func NewProductHandler(l hclog.Logger, c *utils.Configuration, r data.Repository, v *data.Validation) *ProductHandler {
	return &ProductHandler{
		logger: l, 
		config: c, 
		repo: r,
		validator: v,
	}
}

// GenericError is a generic error message returned by server
type GenericError struct {
	Message  string  `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationErrors struct {
	Messages  []string  `json:"messages"`
}

// Used a Key to store the Product Object into context at validation middleware
type ProductKey struct {}

