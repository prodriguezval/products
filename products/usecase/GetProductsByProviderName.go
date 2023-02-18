package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
)

type GetProductsByProviderName struct {
	productProvider provider.ProductProvider
}

func NewGetProductsByProviderNameUseCase(provider provider.ProductProvider) GetProductsByProviderName {
	return GetProductsByProviderName{productProvider: provider}
}

func (u GetProductsByProviderName) Execute(providerName string) []entity.Product {
	providerResult := u.productProvider.GetByProviderName(providerName)
	var response []entity.Product

	if len(providerResult) == 0 {
		return response
	}

	for _, providerProduct := range providerResult {
		response = append(response, providerProduct.ToProduct())
	}
	return response
}
