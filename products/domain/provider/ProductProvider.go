package provider

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	"github.com/prodriguezval/delicaria_products/products/domain/err"
	"github.com/prodriguezval/delicaria_products/products/domain/provider/model"
)

type ProductProvider interface {
	GetAll() ([]model.ProductProviderResponse, *err.ProductProviderError)
	GetByProviderName(providerName string) ([]model.ProductProviderResponse, *err.ProductProviderError)
	Insert(product entity.Product) *err.ProductProviderError
}
