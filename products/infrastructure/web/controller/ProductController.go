package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/web/model"
	"github.com/prodriguezval/delicaria_products/products/usecase"
)

func GetProducts(getAllProducts usecase.GetAllProductsUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		products, err := getAllProducts.Execute()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Message)
		}
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
		products, err := getProductsByProviderName.Execute(providerName)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Message)
		}
		var response []model.ProductResponse
		for _, product := range products {
			productResponse := model.ProductResponse{}
			response = append(response, productResponse.FromProduct(product))
		}
		return ctx.JSON(response)
	}
}

func CreateProduct(createProductUseCase usecase.CreateProduct) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p := new(model.CreateProductRequest)
		err := ctx.BodyParser(p)
		if err != nil {
			return err
		}

		productError := createProductUseCase.Execute(p.ToProduct())
		if productError != nil {
			return fiber.NewError(fiber.StatusInternalServerError, productError.Message)
		}
		return ctx.SendStatus(fiber.StatusCreated)
	}
}
