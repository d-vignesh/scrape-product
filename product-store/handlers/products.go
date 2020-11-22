package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-store/utils"
	"github.com/d-vignesh/scrape-product/product-store/data"
)

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

type GenericError struct {
	Message  string  `json:"message"`
}

type ValidationErrors struct {
	Messages  []string  `json:"messages"`
}

type ProductDetailKey struct {}

