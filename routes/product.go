package routes

import (
	"restapi-with-fiber-and-gorm/database"
	"restapi-with-fiber-and-gorm/models"

	"github.com/gofiber/fiber/v2"
)

// GET http://localhost:3005/api/product
func GetAllProduct(c *fiber.Ctx) error {
	allProduct := []models.Product{}

	result := database.Database.Db.Find(&allProduct)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when getting all data about product",
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting all users data",
		"data": allProduct,
	})
}

// GET http://localhost:3005/api/product/:id
func GetProductById(c *fiber.Ctx) error {
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

	var theProduct models.Product
	
	database.Database.Db.Find(&theProduct, "id = ?", id)
	if theProduct.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The product does not exist",
			"error": "Cannot find product with the given id",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting product by id",
		"data": theProduct,
	})
}

// POST http://localhost:3005/api/product
func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product

	err := c.BodyParser(&newProduct)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	if newProduct.Name == "" || newProduct.SerialNumber == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Please include name (string) and serialNumber (string)",
			"error": "Important field is empthy",
		})
	}

	result := database.Database.Db.Create(&newProduct)
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
		"data": newProduct,
	})
}

// PATCH http://localhost:3005/api/product/:id
func UpdateProduct(c *fiber.Ctx) error {
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

	var theProduct models.Product

	database.Database.Db.Find(&theProduct, "id = ?", id)
	if theProduct.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The product does not exist",
			"error": "Cannot find product with the given id",
		})
	}

	var updatedProduct models.Product

	err = c.BodyParser(&updatedProduct)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	if updatedProduct.Name == "" || updatedProduct.SerialNumber == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Please include name (string) and serialNumber (string)",
			"error": "Important field is empthy",
		})
	}

	theProduct.Name = updatedProduct.Name
	theProduct.SerialNumber = updatedProduct.SerialNumber

	database.Database.Db.Save(&theProduct)

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully update the product",
		"data": theProduct,
	})
}

// DELETE http://localhost:3005/api/product/:id
func DeleteProduct(c *fiber.Ctx) error {
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

	var theProduct models.Product

	database.Database.Db.Find(&theProduct, "id = ?", id)
	if theProduct.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The product does not exist",
			"error": "Cannot find product with the given id",
		})
	}

	err = database.Database.Db.Delete(&theProduct).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Delete operation failed",
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully delete the user",
	}) 
}