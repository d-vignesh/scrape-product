package handlers

import (
	"context"
	"net/http"

	"github.com/d-vignesh/scrape-product/product-store/data"
)

// StoreProduct stores the given product to the database
func (ph *ProductHandler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	
	prod := r.Context().Value(ProductKey{}).(data.Product)

	err := ph.repo.Create(context.Background(), &prod)
	if err != nil {
		ph.logger.Error("unable to insert productdetail into db", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	ph.logger.Info("successfully stored the productDetail")
	w.WriteHeader(http.StatusCreated)
}