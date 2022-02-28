package routes

import "github.com/gofiber/fiber/v2"

// GET http://localhost:3005/api/order
func GetAllOrder(c *fiber.Ctx) error {


	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "",
		"data": "",
	})	
}

// GET http://localhost:3005/api/order/:id
func GetOrderById(c *fiber.Ctx) error {


	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "",
		"data": "",
	})
}

// POST http://localhost:3005/api/order
func CreateOrder(c *fiber.Ctx) error {


	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "",
		"data": "",
	})	
}