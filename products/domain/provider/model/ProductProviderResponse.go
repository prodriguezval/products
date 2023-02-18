package model

import "github.com/prodriguezval/delicaria_products/products/domain/entity"

type ProductProviderResponse struct {
	Id            int
	Name          string
	Provider      string
	BuyPrice      float64
	Tax           float64
	SalePrice     float64
	EarningAmount float64
}

func (p ProductProviderResponse) ToProduct() entity.Product {
	return entity.Product{
		Id:        p.Id,
		Name:      p.Name,
		Provider:  p.Provider,
		BuyPrice:  p.BuyPrice,
		SalePrice: p.SalePrice,
	}
}
