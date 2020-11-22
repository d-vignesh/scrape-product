package handlers

import (
	"context"
	"net/http"

	"github.com/d-vignesh/scrape-product/product-store/data"
)

// Validates the Product json provided in the request body using the Product struct validate tag
func (ph *ProductHandler) MiddlewareValidateProductDetail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ph.logger.Debug("request json", r.Body)
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			ph.logger.Error("unable to deserialize productDetail from request json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		errs := ph.validator.Validate(prod)
		if len(errs) != 0 {
			ph.logger.Error("validation of productDetail request json failed", "error", errs)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&ValidationErrors{Messages: errs.Errors()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), ProductKey{}, *prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}