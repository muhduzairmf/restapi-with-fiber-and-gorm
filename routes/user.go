package routes

import (
	"restapi-with-fiber-and-gorm/database"
	"restapi-with-fiber-and-gorm/models"

	"github.com/gofiber/fiber/v2"
)

// GET http://localhost:3005/api/user
func GetAllUser(c *fiber.Ctx) error {
	allUsers := []models.User{}

	result := database.Database.Db.Find(&allUsers)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when getting all data about users",
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting all users data",
		"data": allUsers,
	})
}

// GET http://localhost:3005/api/user/:id
func GetUserById(c *fiber.Ctx) error {
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

	var theUser models.User

	database.Database.Db.Find(&theUser, "id = ?", id)
	if theUser.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully getting the user by id",
		"data": theUser,
	})
}

// POST http://localhost:3005/api/user
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

// PATCH http://localhost:3005/api/user/:id
func UpdateUser(c *fiber.Ctx) error {
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

	var theUser models.User

	database.Database.Db.Find(&theUser, "id = ?", id)
	if theUser.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	var updatedUser models.User

	err = c.BodyParser(&updatedUser)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		}) 
	}

	if updatedUser.FirstName == "" || updatedUser.LastName == "" || updatedUser.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include firstName (string), lastName (string), and email (string)",
		})
	}

	theUser.FirstName = updatedUser.FirstName
	theUser.LastName = updatedUser.LastName
	theUser.Email = updatedUser.Email

	database.Database.Db.Save(&theUser)

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully update the user",
		"data": theUser,
	}) 
}

// DELETE http://localhost:3005/api/user/:id
func DeleteUser(c *fiber.Ctx) error {
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

	var theUser models.User

	database.Database.Db.Find(&theUser, "id = ?", id)
	if theUser.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	err = database.Database.Db.Delete(&theUser).Error
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
