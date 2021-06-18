package use_case

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/graph/model"
	"github.com/aditiapratama1231/graphql-example/internal/product/repository"
)

// product service
type Product struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewProductResolver(productRepo repository.ProductRepositoryInterface) Product {
	return Product{
		ProductRepository: productRepo,
	}
}

func (p Product) GetProducts(ctx context.Context) []*model.Product {
	return p.ProductRepository.GetProducts(ctx)
}

func (p Product) GetSingleProduct(ctx context.Context) *model.Product {
	return p.ProductRepository.GetSingleProduct(ctx)
}
