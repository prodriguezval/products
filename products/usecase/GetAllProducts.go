package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
)

type GetAllProductsUseCase struct {
	productProvider provider.ProductProvider
}

func NewGetAllProductsUseCase(provider provider.ProductProvider) GetAllProductsUseCase {
	return GetAllProductsUseCase{productProvider: provider}
}

func (u GetAllProductsUseCase) Execute() []entity.Product {
	providerResult := u.productProvider.GetAll()
	var response []entity.Product

	if len(providerResult) == 0 {
		return response
	}

	for _, providerProduct := range providerResult {
		response = append(response, providerProduct.ToProduct())
	}
	return response
}
