package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/provider"
	"github.com/prodriguezval/delicaria_products/products/infrastructure/web/controller"
	"log"
)

func main() {
	app := fiber.New()
	productProvider := provider.NewSheetProductProvider()
	products := app.Group("/products")
	controller.ProductController(products, productProvider)
	log.Fatal(app.Listen(":3000"))
}
