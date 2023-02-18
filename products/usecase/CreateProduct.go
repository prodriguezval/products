package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
	"log"
)

type CreateProduct struct {
	productProvider provider.ProductProvider
}

func NewCreateProduct(provider provider.ProductProvider) CreateProduct {
	return CreateProduct{productProvider: provider}
}

func (u CreateProduct) Execute(product entity.Product) {
	log.Printf("Trying to save product %v", product)
	u.productProvider.Insert(product)
}
