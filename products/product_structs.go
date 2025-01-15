package products

import (
	"encoding/json"
	"fmt"
	"os"
)

// Variable to storage products data
var ProductsData []Product

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func LoadProductsFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file failed: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ProductsData)

	if err != nil {
		return fmt.Errorf("decoude products failed: %s", err)
	}

	return nil
}
