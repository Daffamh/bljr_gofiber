package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetStatus(c *fiber.Ctx) error {
	var status []models.Status
	config.DB.Unscoped().
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Find(&status)
	return c.JSON(status)
}

func GetStatuss(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	if err := config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
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
	status.CreatedById = &userID
	status.UpdatedById = &userID

	if err := config.DB.Create(&status).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStatus models.Status
	if err := config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
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
	status.UpdatedById = &userID

	config.DB.Save(&status)

	var updatedStatus models.Status
	config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Where("id = ?", id).First(&updatedStatus)

	return c.JSON(updatedStatus)
}
func DeleteStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	if err := config.DB.Delete(&models.Status{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete status",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.Status{}).Where("id = ?", id).Update("deleted_by_id", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update delete_by field",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "status deletd successfully",
	})
}
