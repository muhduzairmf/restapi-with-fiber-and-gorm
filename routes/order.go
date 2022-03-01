package routes

import (
	"restapi-with-fiber-and-gorm/database"
	"restapi-with-fiber-and-gorm/models"

	"github.com/gofiber/fiber/v2"
)

// GET http://localhost:3005/api/order
func GetAllOrder(c *fiber.Ctx) error {
	allOrder := []models.Order{}

	result := database.Database.Db.Find(&allOrder)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when getting all data about product",
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting all order data",
		"data": allOrder,
	})	
}

// GET http://localhost:3005/api/order/:id
func GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be an integer",
			"error": err.Error(),
		})
	}

	if id < 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be a positive integer",
			"error": "The id params has value less than zero (0).",
		})
	} 
	
	var theOrder models.Order

	database.Database.Db.Find(&theOrder, "id = ?", id)
	if theOrder.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The order does not exist",
			"error": "Cannot find order with the given id",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting order by id",
		"data": theOrder,
	})
}

// POST http://localhost:3005/api/order
func CreateOrder(c *fiber.Ctx) error {
	var newOrder models.Order

	err := c.BodyParser(&newOrder)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	if newOrder.ProductId == 0 || newOrder.AmountToOrder == 0 || newOrder.UserId == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Please include productId (integer), amountToOrder (integer) and userId (integer)",
			"error": "Important field is empthy",
		})
	}

	result := database.Database.Db.Create(&newOrder)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when creating data",
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully create a new product",
		"data": newOrder,
	})	
}