// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Brand struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
}

type CreateProductInput struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type Product struct {
	ID      string  `json:"id"`
	Name    *string `json:"name"`
	Stock   *int    `json:"stock"`
	BrandID *string `json:"brand_id"`
}
