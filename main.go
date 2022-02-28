package main

import (
	"log"
	"restapi-with-fiber-and-gorm/database"
	"restapi-with-fiber-and-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func userRoute(app *fiber.App) {
	app.Get("/api/user", routes.GetAllUser)
	app.Get("/api/user/:id", routes.GetUserById)
	app.Post("/api/user", routes.CreateUser)
	app.Patch("/api/user/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
}

func productRoute(app *fiber.App) {
	app.Get("/api/product", routes.GetAllProduct)
	app.Get("/api/product/:id", routes.GetProductById)
	app.Post("/api/product", routes.CreateProduct)
	app.Patch("/api/product/:id", routes.UpdateProduct)
	app.Delete("/api/product/:id", routes.DeleteProduct)
}

func orderRoute(app *fiber.App) {
	app.Get("/api/order", routes.GetAllOrder)
	app.Get("/api/order/:id", routes.GetOrderById)
	app.Post("/api/order", routes.CreateOrder)
}

func notFoundRoute(app *fiber.App) {
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status": "Not found",
			"message": "Routes not found",
		})
	})
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "Ok",
			"message": "Welcome to home page!",
		})
	})

	userRoute(app)

	productRoute(app)

	orderRoute(app)

	notFoundRoute(app)

	log.Println(app.Listen(":3005"))
}
