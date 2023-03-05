package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/provider"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/web/controller"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/web/middleware"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	productProvider := provider.NewSheetProductProvider()
	products := app.Group("/products")
	controller.ProductController(products, productProvider)
	log.Fatal(app.Listen(":3000"))
}
