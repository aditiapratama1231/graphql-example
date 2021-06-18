package graph

import (
	"github.com/aditiapratama1231/graphql-example/graph/dataloader"
	"github.com/aditiapratama1231/graphql-example/internal/brand"
	"github.com/aditiapratama1231/graphql-example/internal/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductResolver product.ProductInterface
	BrandResolver   brand.BrandInterface
	DataLoaders     dataloader.Retriever
}
