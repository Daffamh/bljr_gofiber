package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

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
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Ganti dengan delete dari database " + id,
	})
}
