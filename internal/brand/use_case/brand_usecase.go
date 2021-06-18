package use_case

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/graph/model"
	"github.com/aditiapratama1231/graphql-example/internal/brand/repository"
)

// brand service
type Brand struct {
	BrandRepository repository.BrandRepositoryInterface
}

func NewProductResolver(productRepo repository.BrandRepositoryInterface) Brand {
	return Brand{
		BrandRepository: productRepo,
	}
}

func (b Brand) GetBrands(ctx context.Context) []*model.Brand {
	return b.BrandRepository.GetBrands(ctx)
}
