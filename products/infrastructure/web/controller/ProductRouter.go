package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prodriguezval/delicaria_products/products/domain/provider"
	"github.com/prodriguezval/delicaria_products/products/usecase"
)

func ProductController(app fiber.Router, productProvider provider.ProductProvider) {
	app.Get("/", GetProducts(usecase.NewGetAllProductsUseCase(productProvider)))
	app.Get("/provider/:provider_name", GetProductsByProviderName(usecase.NewGetProductsByProviderNameUseCase(productProvider)))
	app.Post("/", CreateProduct(usecase.NewCreateProduct(productProvider)))
}
