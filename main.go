package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/melisebestrada/go-web-api/products"
)

func main() {
	products, err := getProductsData()
	if err != nil {
		fmt.Println("error getting products data: ", err)
		return
	}

	fmt.Println(products)
}

func getProductsData() ([]products.Product, error) {
	file, err := os.Open("data/products.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []products.Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}
