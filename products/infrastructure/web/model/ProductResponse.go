package model

import "github.com/prodriguezval/delicaria_products/products/domain/entity"

type ProductResponse struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Provider      string  `json:"provider"`
	BuyPrice      float64 `json:"buy_price"`
	TaxAmount     float64 `json:"tax_amount"`
	EarningAmount float64 `json:"earning_amount"`
	SubTotal      float64 `json:"sub_total"`
	Total         float64 `json:"total"`
}

func (p ProductResponse) FromProduct(product entity.Product) ProductResponse {
	return ProductResponse{
		Id:            product.Id,
		Name:          product.Name,
		Provider:      product.Provider,
		EarningAmount: product.CalculateEarningAmount(),
		TaxAmount:     product.CalculateTaxAmount(),
		BuyPrice:      product.BuyPrice,
		SubTotal:      product.CalculateSubTotal(),
		Total:         product.SalePrice,
	}
}
