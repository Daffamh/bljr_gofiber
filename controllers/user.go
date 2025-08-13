package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"
	"gorm.io/gorm/clause"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"id": id,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())

	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}
	config.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}
	return c.JSON(fiber.Map{
		"id": id,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	user.DeletedBy = &userID

	if err := config.DB.Clauses(clause.Returning{}).Save(&user).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
