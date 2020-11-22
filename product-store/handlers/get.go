package handlers

import (
	"context"
	"net/http"

	"github.com/d-vignesh/scrape-product/product-store/data"
)

// Gets list of all products
func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	prods, err := ph.repo.GetAll(context.Background())
	if err != nil {
		ph.logger.Error("unable to read products from DB", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	data.ToJSON(prods, w)
	w.WriteHeader(http.StatusOK)
}