package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(ProductsData)
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	idProduct, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Enter a valid id", http.StatusBadRequest)
		return
	}

	for _, prod := range ProductsData {
		if prod.Id == idProduct {
			w.Header().Add("content-type", "application/json")
			json.NewEncoder(w).Encode(prod)
			return
		}
	}

	http.Error(w, "The id does not exist", http.StatusNotFound)

}

func GetProductByPriceGt(w http.ResponseWriter, r *http.Request) {
	price := r.URL.Query().Get("priceGt")
	if price == "" {
		http.Error(w, "priceGt param is required", http.StatusBadRequest)
		return
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Enter a valid price", http.StatusBadRequest)
		return
	}

	var products []Product

	for _, prod := range ProductsData {
		if prod.Price > priceFloat {
			products = append(products, prod)
		}
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}
