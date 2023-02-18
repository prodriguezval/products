package entity

import "math"

type Product struct {
	Id        int
	Name      string
	Provider  string
	BuyPrice  float64
	SalePrice float64
}

var taxPercentage float64 = 19

func (p Product) CalculateTaxAmount() float64 {
	taxAmount := p.CalculateSubTotal() * (taxPercentage / 100)
	return math.Round(taxAmount*100) / 100
}

func (p Product) CalculateSubTotal() float64 {
	subTotal := p.SalePrice / (1 + (taxPercentage / 100))
	return math.Round(subTotal*100) / 100
}

func (p Product) CalculateEarningAmount() float64 {
	earningAmount := p.SalePrice - p.BuyPrice - p.CalculateTaxAmount()
	return math.Round(earningAmount*100) / 100
}


