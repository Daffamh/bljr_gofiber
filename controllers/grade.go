package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetGrade(c *fiber.Ctx) error {
	var grade []models.Grade
	config.DB.Unscoped().
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		Find(&grade)
	return c.JSON(grade)
}

func GetGrades(c *fiber.Ctx) error {
	id := c.Params("id")
	var grades models.Grade
	if err := config.DB.
		Preload("UpdatedBy").
		Preload("CreatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		First(&grades, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Grades not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(grades)
}

func CreateGrade(c *fiber.Ctx) error {
	var grade models.Grade
	validate := validator.New()

	err := c.BodyParser(&grade)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid grade",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(grade)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("id").(uint)
	grade.CreatedById = &userID
	grade.UpdatedById = &userID

	if err := config.DB.Create(&grade).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdGrade models.Grade
	if err := config.DB.
		Preload("UpdatedBy").
		Preload("CreatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		First(&createdGrade, grade.Id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(createdGrade)
}

func UpdateGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	var grade models.Grade
	if err := config.DB.First(&grade, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "grade not found",
			"error":   err.Error(),
		})

	}
	if err := c.BodyParser(&grade); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}

	if err := c.BodyParser(&grade); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("id").(uint)
	grade.UpdatedById = &userID

	config.DB.Save(&grade)

	var updatedGrade models.Grade
	config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		Where("id = ?", id).First(&updatedGrade)

	return c.JSON(updatedGrade)
}

func DeleteGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	if err := config.DB.Delete(&models.Grade{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete grade",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.Grade{}).Where("id = ?", id).Update("deleted_by_id", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update delete_by field",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "grade deletd successfully",
	})
}
