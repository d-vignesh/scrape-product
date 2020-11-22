package handlers

import (
	"net/http"

	"github.com/d-vignesh/scrape-product/product-scraper/data"
)

func (sh *ScrapingHandler) ScrapeURL(w http.ResponseWriter, r *http.Request) {
	
	url := r.URL.Query().Get("url")
	if url == "" {
		sh.logger.Error("url not provided in the request")
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&GenericError{Message: "the parameter 'url' to be scraped is required and not provided"}, w)
		return
	}

	prod := sh.scraper.ScrapeURL(url)
	w.WriteHeader(http.StatusOK)
	data.ToJSON(prod, w)
}