package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetHomeRoom(c *fiber.Ctx) error {
	var homeroom []models.HomeRoom
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		Find(&homeroom)
	return c.JSON(homeroom)
}

func GetHomeRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	var homeroom models.HomeRoom
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		First(&homeroom, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Home not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(homeroom)
}

func CreateHomeRoom(c *fiber.Ctx) error {
	var homeroom models.HomeRoom
	validate := validator.New()

	err := c.BodyParser(&homeroom)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid homeroom",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(homeroom)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("id").(uint)
	homeroom.CreatedById = userID
	homeroom.UpdatedById = userID

	if err := config.DB.Create(&homeroom).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdHomeRoom models.HomeRoom
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		First(&createdHomeRoom, homeroom.Id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "homeroom not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(createdHomeRoom)
}

func UpdateHomeRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var homeroom models.HomeRoom
	if err := config.DB.First(&homeroom, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "homeroom not found",
			"error":   err.Error(),
		})

	}
	if err := c.BodyParser(&homeroom); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}

	if err := c.BodyParser(&homeroom); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("id").(uint)
	homeroom.UpdatedById = userID

	config.DB.Save(&homeroom)

	var updatedHomeRoom models.Status
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		Where("id = ?", id).First(&updatedHomeRoom)

	return c.JSON(updatedHomeRoom)
}
func DeleteHomeRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	if err := config.DB.Delete(&models.HomeRoom{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete home room",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.HomeRoom{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update delete_by field",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "home room deletd successfully",
	})
}
