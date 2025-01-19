package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/melisebestrada/go-web-api/internal/domain"
)

type ProductsRepositoryInterface interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id int, product domain.Product) (domain.Product, error)
	PatchProduct(id int, product domain.Product) (domain.Product, error)
}

type productsRepository struct {
	products []domain.Product
}

func NewProductRepository(filePath string) (ProductsRepositoryInterface, error) {
	repository := &productsRepository{}

	err := repository.loadProductsFromFile(filePath)
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (pr *productsRepository) GetAllProducts() ([]domain.Product, error) {
	if pr.products == nil {
		return nil, fmt.Errorf("no products found")
	}

	return pr.products, nil
}

func (pr *productsRepository) GetProductById(id int) (domain.Product, error) {

	for _, product := range pr.products {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("product does not exist")

}

func (pr *productsRepository) loadProductsFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pr.products); err != nil {
		return err
	}

	return nil
}

func (pr *productsRepository) CreateProduct(product domain.Product) (domain.Product, error) {
	for _, prod := range pr.products {
		if prod.CodeValue == product.CodeValue {
			return domain.Product{}, fmt.Errorf("code value already exists")
		}
	}

	var newProduct = domain.Product{
		Id:          len(pr.products) + 1,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	pr.products = append(pr.products, newProduct)

	return newProduct, nil

}

func (pr *productsRepository) UpdateProduct(id int, product domain.Product) (domain.Product, error) {
	oldProduct := -1
	for index, prod := range pr.products {
		if prod.CodeValue == product.CodeValue && id != prod.Id {
			return domain.Product{}, fmt.Errorf("code value already exists")
		}

		if id == prod.Id {
			oldProduct = index
		}
	}

	if oldProduct == -1 {
		return domain.Product{}, fmt.Errorf("product with id %d not found", id)
	}

	product.Id = id
	pr.products[oldProduct] = product

	return pr.products[oldProduct], nil
}

func (pr *productsRepository) PatchProduct(id int, product domain.Product) (domain.Product, error) {
	var productIndex = -1
	for index, prod := range pr.products {
		if prod.CodeValue == product.CodeValue && id != prod.Id {
			return domain.Product{}, fmt.Errorf("code value already exists")
		}

		if id == prod.Id {
			productIndex = index
		}
	}

	if productIndex == -1 {
		return domain.Product{}, fmt.Errorf("product with id %d not found", id)
	}

	if product.Name != "" {
		pr.products[productIndex].Name = product.Name
	}
	if product.Quantity > 0 {
		pr.products[productIndex].Quantity = product.Quantity
	}
	if product.CodeValue != "" {
		pr.products[productIndex].CodeValue = product.CodeValue
	}
	if product.IsPublished != pr.products[productIndex].IsPublished {
		pr.products[productIndex].IsPublished = product.IsPublished
	}
	if product.Expiration != "" {
		pr.products[productIndex].Expiration = product.Expiration
	}
	if product.Price > 0 {
		pr.products[productIndex].Price = product.Price
	}

	return pr.products[productIndex], nil
}
