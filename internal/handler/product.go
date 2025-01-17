package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/melisebestrada/go-web-api/internal/service"
)

type ProductHandler struct {
	service service.ProductServiceInterface
}

func NewProductHandler(service service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := ph.service.GetAllProducts()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (ph *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idProduct, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Enter a valid id", http.StatusBadRequest)
			return
		}

		product, err := ph.service.GetProductById(idProduct)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)

	}
}

func (ph *ProductHandler) SearchPriceGt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		price := r.URL.Query().Get("priceGt")

		if price == "" {
			http.Error(w, "price is mandatory", 400)
			return
		}

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "price is mandatory", 400)
			return
		}

		products, err := ph.service.SearchPriceGt(priceFloat)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}
