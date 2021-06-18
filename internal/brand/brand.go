package brand

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/graph/model"
)

type BrandInterface interface {
	GetBrands(ctx context.Context) []*model.Brand
}
