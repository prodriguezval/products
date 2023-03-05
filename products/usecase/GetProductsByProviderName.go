package usecase

import (
	"github.com/prodriguezval/delicaria_products/products/domain/entity"
	domainError "github.com/prodriguezval/delicaria_products/products/domain/err"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
	"log"
)

type GetProductsByProviderName struct {
	productProvider provider.ProductProvider
}

func NewGetProductsByProviderNameUseCase(provider provider.ProductProvider) GetProductsByProviderName {
	return GetProductsByProviderName{productProvider: provider}
}

func (u GetProductsByProviderName) Execute(providerName string) ([]entity.Product, *domainError.ProductProviderError) {
	providerResult, err := u.productProvider.GetByProviderName(providerName)
	if err != nil {
		log.Println(err.Error())
		return nil, domainError.NewProductProviderError("Error executing GetProductsByProviderName use case", err)
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
