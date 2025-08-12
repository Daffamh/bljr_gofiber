package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetStatus(c *fiber.Ctx) error {
	var status []models.Status
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Find(&status)
	return c.JSON(status)
}

func GetStatuss(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		First(&status, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "status not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(status)
}

func CreateStatus(c *fiber.Ctx) error {
	var status models.Status
	validate := validator.New()

	err := c.BodyParser(&status)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid status",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(status)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("id").(uint)
	status.CreatedBy = userID
	status.UpdatedBy = userID

	if err := config.DB.Create(&status).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStatus models.Status
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		First(&createdStatus, status.Id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "status not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(createdStatus)
}

func UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	if err := config.DB.First(&status, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "status not found",
			"error":   err.Error(),
		})

	}
	if err := c.BodyParser(&status); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}

	if err := c.BodyParser(&status); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("id").(uint)
	status.UpdatedBy = userID

	config.DB.Save(&status)

	var updatedHomeRoom models.Status
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Where("id = ?", id).First(&updatedHomeRoom)

	return c.JSON(updatedHomeRoom)
}
func DeleteStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	if err := config.DB.Delete(&status, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete status",
			"error":   err.Error(),
		})

	}
	return c.JSON(fiber.Map{
		"message": "status deleted successfully",
	})
}
