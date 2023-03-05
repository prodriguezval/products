package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	domainError "github.com/prodriguezval/delicaria_products/products/domain/err"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
)

type GetAllProductsUseCase struct {
	productProvider provider.ProductProvider
}

func NewGetAllProductsUseCase(provider provider.ProductProvider) GetAllProductsUseCase {
	return GetAllProductsUseCase{productProvider: provider}
}

func (u GetAllProductsUseCase) Execute() ([]entity.Product, *domainError.ProductProviderError) {
	providerResult, err := u.productProvider.GetAll()
	if err != nil {
		return nil, domainError.NewProductProviderError("Error executing Create GetAllProductsUseCase use case", err)
	}
	var response []entity.Product

	if len(providerResult) == 0 {
		return response, nil
	}

	for _, providerProduct := range providerResult {
		response = append(response, providerProduct.ToProduct())
	}
	return response, nil
}
