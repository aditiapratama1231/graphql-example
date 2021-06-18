package product

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/graph/model"
)

type ProductInterface interface {
	GetProducts(ctx context.Context) []*model.Product
	GetSingleProduct(ctx context.Context) *model.Product
}
