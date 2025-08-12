package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetStudentGrade(c *fiber.Ctx) error {
	var studentgrade []models.StudentGrade
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("Student").
		Preload("Grade").
		Preload("HomeRoom").
		Preload("Status").
		Find(&studentgrade)
	return c.JSON(studentgrade)
}

func GetStudentGrades(c *fiber.Ctx) error {
	id := c.Params("id")
	var studentgrade models.StudentGrade
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		First(&studentgrade, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "studentgrade not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(studentgrade)
}

func CreateStudentGrade(c *fiber.Ctx) error {
	var studentgrade models.StudentGrade
	//validate := validator.New()

	err := c.BodyParser(&studentgrade)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid studentgrade",
			"error":   err.Error(),
		})
	}

	//err = validate.Struct(studentgrade)
	//if err != nil {
	//	return c.Status(400).JSON(fiber.Map{
	//		"success": false,
	//		"message": "Invalid data request",
	//		"error":   err.Error(),
	//	})
	//}

	userID := c.Locals("id").(uint)
	studentgrade.CreatedBy = userID
	studentgrade.UpdatedBy = userID

	if err := config.DB.Create(&studentgrade).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStudentGrade models.StudentGrade
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("Student").
		Preload("Grade").
		Preload("HomeRoom").
		Preload("Status").
		First(&createdStudentGrade, studentgrade.Id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "studentgrade not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(createdStudentGrade)
}

func UpdateStudentGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	var studentGrade models.StudentGrade
	if err := config.DB.First(&studentGrade, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "studentGrade not found",
			"error":   err.Error(),
		})

	}
	if err := c.BodyParser(&studentGrade); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}

	userID := c.Locals("id").(uint)
	studentGrade.UpdatedBy = userID

	config.DB.Save(&studentGrade)

	var updatedStudentGrade models.StudentGrade
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("Student").
		Preload("Grade").
		Preload("HomeRoom").
		Preload("Status").
		Where("id = ?", id).First(&updatedStudentGrade)

	return c.JSON(updatedStudentGrade)
}
func DeleteStudentGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	var studentgrade models.StudentGrade
	if err := config.DB.Delete(&studentgrade, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete studentgrade",
			"error":   err.Error(),
		})

	}
	return c.JSON(fiber.Map{
		"message": "studentgrade deleted successfully",
	})
}
