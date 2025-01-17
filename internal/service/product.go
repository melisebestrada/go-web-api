package service

import (
	"fmt"

	"github.com/melisebestrada/go-web-api/internal/domain"
	"github.com/melisebestrada/go-web-api/internal/repository"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
}

type productService struct {
	repository repository.ProductsRepositoryInterface
}

func NewProductService(repository repository.ProductsRepositoryInterface) ProductServiceInterface {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	products, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	return products, nil
}

func (ps *productService) GetProductById(id int) (domain.Product, error) {
	product, err := ps.repository.GetProductById(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, err
}

func (ps *productService) SearchPriceGt(price float64) ([]domain.Product, error) {
	products, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	var productsGt []domain.Product

	for _, product := range products {
		if product.Price > price {
			productsGt = append(productsGt, product)
		}
	}
	return productsGt, nil
}
