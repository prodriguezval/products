package model

import "github.com/prodriguezval/delicaria_products/products/domain/entity"

type CreateProductRequest struct {
	Name         string  `json:"name"`
	ProviderName string  `json:"provider_name"`
	BuyPrice     float64 `json:"buy_price"`
	SalePrice    float64 `json:"sale_price"`
}

func (p CreateProductRequest) ToProduct() entity.Product {
	return entity.Product{
		Name:      p.Name,
		Provider:  p.ProviderName,
		BuyPrice:  p.BuyPrice,
		SalePrice: p.SalePrice,
	}
}
