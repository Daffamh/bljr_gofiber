package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/config"
	"gofiberapp/models"
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

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.User{}).Where("id = ?", id).Update("deleted_by_id", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update delete_by field",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "user deletd successfully",
	})
}
