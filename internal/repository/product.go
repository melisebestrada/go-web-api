package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/melisebestrada/go-web-api/internal/domain"
	"github.com/melisebestrada/go-web-api/pkg/web"
)

type ProductsRepositoryInterface interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	CreateProduct(product web.RequestBodyProduct) (domain.Product, error)
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

func (pr *productsRepository) CreateProduct(product web.RequestBodyProduct) (domain.Product, error) {
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
