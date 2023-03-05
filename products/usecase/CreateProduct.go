package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	domainError "github.com/prodriguezval/delicaria_products/products/domain/err"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
	"log"
)

type CreateProduct struct {
	productProvider provider.ProductProvider
}

func NewCreateProduct(provider provider.ProductProvider) CreateProduct {
	return CreateProduct{productProvider: provider}
}

func (u CreateProduct) Execute(product entity.Product) *domainError.ProductProviderError {
	log.Printf("Trying to save product %v", product)
	err := u.productProvider.Insert(product)
	if err != nil {
		return domainError.NewProductProviderError("Error executing Create CreateProduct use case", err)
	}
	return nil
}
