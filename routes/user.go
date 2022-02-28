package routes

import (
	"restapi-with-fiber-and-gorm/database"
	"restapi-with-fiber-and-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting all users data",
		"data": "hello world!",
	})
}

func CreateUser(c *fiber.Ctx) error {
	var newUser models.User
	
	
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}
	
	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include firstName (string), lastName (string), and email (string)",
		})
	}
	
	result := database.Database.Db.Create(&newUser)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when creating data",
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully create a new user",
		"data": newUser,
	})
	
}