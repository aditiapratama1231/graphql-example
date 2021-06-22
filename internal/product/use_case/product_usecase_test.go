package use_case

import (
	"context"
	"testing"

	"github.com/aditiapratama1231/graphql-example/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r ProductRepositoryMock) GetProducts(ctx context.Context) []*model.Product {
	name := "Lg"
	stock := 10
	products := []*model.Product{
		&model.Product{
			ID:    "1",
			Name:  &name,
			Stock: &stock,
		},
		&model.Product{
			ID:    "2",
			Name:  &name,
			Stock: &stock,
		},
	}

	return products
}

func (r ProductRepositoryMock) GetSingleProduct(ctx context.Context) *model.Product {
	name := "Lg"
	stock := 10
	products := model.Product{
		ID:    "1",
		Name:  &name,
		Stock: &stock,
	}

	return &products
}

func TestProduct_GetProducts(t *testing.T) {
	ctx := context.Background()
	name_want := "Lg"

	count_want := 2

	repo := ProductRepositoryMock{}
	repo.On("GetProducts").Return([]*model.Product{})

	product_service := Product{repo}
	products := product_service.GetProducts(ctx)

	for i := range products {
		assert.Equal(t, products[i].Name, &name_want, "Name must be LG")
	}

	if len(products) != count_want {
		t.Errorf("Products count return %d, expected %d", len(products), count_want)
	}

}
