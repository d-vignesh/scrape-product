package handlers

import (
	"context"
	"net/http"

	"github.com/d-vignesh/scrape-product/product-store/data"
)

func (ph *ProductHandler) Store(w http.ResponseWriter, r *http.Request) {
	
	prodDetail := r.Context().Value(ProductDetailKey{}).(data.ProductDetail)

	err := ph.repo.Create(context.Background(), &prodDetail)
	if err != nil {
		ph.logger.Error("unable to insert productdetail into db", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	ph.logger.Info("successfully stored the productDetail")
	w.WriteHeader(http.StatusCreated)
}