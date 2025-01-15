package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/melisebestrada/go-web-api/products"
)

func main() {
	err := products.LoadProductsFromFile("data/products.json")
	if err != nil {
		fmt.Println("error getting products data: ", err)
		return
	}

	r := chi.NewRouter()

	r.Get("/ping", getPong)

	// product router
	r.Route("/product", func(r chi.Router) {
		r.Get("/", products.GetAllProducts)
		r.Get("/{id}", products.GetProductById)
	})

	fmt.Println("Server started at http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}

func getPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
