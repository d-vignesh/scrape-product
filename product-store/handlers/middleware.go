package handlers

import (
	"context"
	"net/http"

	"github.com/d-vignesh/scrape-product/product-store/data"
)

func (ph *ProductHandler) MiddlewareValidateProductDetail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ph.logger.Debug("request json", r.Body)
		prodDetail := &data.ProductDetail{}

		err := data.FromJSON(prodDetail, r.Body)
		if err != nil {
			ph.logger.Error("unable to deserialize productDetail from request json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		errs := ph.validator.Validate(prodDetail)
		if len(errs) != 0 {
			ph.logger.Error("validation of productDetail request json failed", "error", errs)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&ValidationErrors{Messages: errs.Errors()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), ProductDetailKey{}, *prodDetail)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}