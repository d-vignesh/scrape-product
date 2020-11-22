package handlers

import (
	"net/http"
	"bytes"
	"encoding/json"

	"github.com/d-vignesh/scrape-product/product-scraper/data"
)

// ScrapeURL scrapes the url provided in the request param
func (sh *ScrapingHandler) ScrapeURL(w http.ResponseWriter, r *http.Request) {
	
	// checking for url parameter in the request
	url := r.URL.Query().Get("url")
	if url == "" {
		sh.logger.Error("url not provided in the request")
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&GenericError{Message: "the parameter 'url' to be scraped is required and not provided"}, w)
		return
	}

	// scraping the given url using scraper service
	prod := sh.scraper.ScrapeURL(url)
	prod.URL = url
	// prodDetail := &data.ProductDetail{
	// 	URL: url,
	// 	Product: *prod,
	// }

	// marshalling the product and passing it to Product Store service 
	requestBody, err := json.Marshal(*prod)
	if err != nil {
		sh.logger.Error("unable to marshal product detail", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: "unable to store the product detail, failed to marshal the productDetail"}, w)
		return
	}

	sh.logger.Debug("sending post request to product store at", sh.config.StoreServerAddress)

	resp, err := http.Post(sh.config.StoreServerAddress, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		sh.logger.Error("unable to persist the product detail", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: "unable to store the product detail, request to ProductStore failed"}, w)
		return
	}

	sh.logger.Info("store product status", resp.Status)
	w.WriteHeader(http.StatusOK)
	data.ToJSON(prod, w)
}