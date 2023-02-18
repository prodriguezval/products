package provider

import "github.com/prodriguezval/delicaria_products/products/domain/provider/model"

type ProductProvider interface {
	GetAll() []model.ProductProviderResponse
	GetByProviderName(providerName string) []model.ProductProviderResponse
}
