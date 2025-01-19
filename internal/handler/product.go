package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/melisebestrada/go-web-api/internal/domain"
	"github.com/melisebestrada/go-web-api/internal/service"
	"github.com/melisebestrada/go-web-api/pkg/validations"
	"github.com/melisebestrada/go-web-api/pkg/web"
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

func (ph *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody domain.Product
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.SendResponse(w, "Bad Request", nil, true, http.StatusBadRequest)
			return
		}

		validationFail := validations.ValidatedEmptyFields(w, reqBody)
		if validationFail {
			return
		}

		err := validations.ValidateDate(reqBody.Expiration)
		if err != nil {
			web.SendResponse(w, err.Error(), nil, true, http.StatusBadRequest)
			return
		}

		newProduct, err := ph.service.CreateProduct(reqBody)
		if err != nil {
			web.SendResponse(w, err.Error(), nil, true, http.StatusBadRequest)
			return
		}

		web.SendResponse(w, "Product created", newProduct, false, http.StatusCreated)

	}
}

func (ph *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idProduct, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Enter a valid id", http.StatusBadRequest)
			return
		}

		var reqBody domain.Product
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.SendResponse(w, "Bad request", nil, true, http.StatusBadRequest)
			return
		}

		validationFail := validations.ValidatedEmptyFields(w, reqBody)
		if validationFail {
			return
		}

		err = validations.ValidateDate(reqBody.Expiration)
		if err != nil {
			web.SendResponse(w, err.Error(), nil, true, http.StatusBadRequest)
			return
		}

		updatedProduct, err := ph.service.UpdateProduct(idProduct, reqBody)
		if err != nil {
			web.SendResponse(w, err.Error(), nil, true, http.StatusBadRequest)
			return
		}

		web.SendResponse(w, "Product updated", updatedProduct, false, http.StatusOK)

	}
}
