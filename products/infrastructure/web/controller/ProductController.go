package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/web/model"
	"github.com/prodriguezval/delicaria_products/products/usecase"
)

func GetProducts(getAllProducts usecase.GetAllProductsUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		products := getAllProducts.Execute()
		var response []model.ProductResponse
		for _, product := range products {
			productResponse := model.ProductResponse{}
			response = append(response, productResponse.FromProduct(product))
		}
		return ctx.JSON(response)
	}
}

func GetProductsByProviderName(getProductsByProviderName usecase.GetProductsByProviderName) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		providerName := ctx.Params("provider_name")
		products := getProductsByProviderName.Execute(providerName)
		var response []model.ProductResponse
		for _, product := range products {
			productResponse := model.ProductResponse{}
			response = append(response, productResponse.FromProduct(product))
		}
		return ctx.JSON(response)
	}
}
