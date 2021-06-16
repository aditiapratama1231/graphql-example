package product

import (
	"github.com/aditiapratama1231/graphql-example/graph/model"
	"github.com/aditiapratama1231/graphql-example/internal/repository"
)

type Product struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewProductResolver(productRepo repository.ProductRepositoryInterface) Product {
	return Product{
		ProductRepository: productRepo,
	}
}

type ProductInterface interface {
	GetProducts() []model.Product
}

func (p Product) GetProducts() []*model.Product {
	return p.ProductRepository.GetProducts()
}
