package repository

import (
	"github.com/aditiapratama1231/graphql-example/graph/model"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts() []*model.Product
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (pr ProductRepository) GetProducts() []*model.Product {
	var products []*model.Product

	pr.db.Find(&products)

	return products
}
