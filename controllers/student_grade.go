package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/gofiber/fiber/v2"
)

func GetStudentGrade(c *fiber.Ctx) error {
	var studentgrade []models.StudentGrade
	config.DB.
		Preload("UpdatedBy").
		Preload("CreatedBy").
		Preload("DeletedBy").
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
		Preload("UpdatedBy").
		Preload("CreatedBy").
		Preload("DeletedBy").
		Preload("Student").
		Preload("Grade").
		Preload("HomeRoom").
		Preload("Status").
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
	studentgrade.CreatedById = &userID
	studentgrade.UpdatedById = &userID

	if err := config.DB.Create(&studentgrade).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStudentGrade models.StudentGrade
	if err := config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
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
	studentGrade.UpdatedById = &userID

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
	userID := c.Locals("id").(uint)

	if err := config.DB.Delete(&models.StudentGrade{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete student field",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.StudentGrade{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update delete_by student grade",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "student grade delete successfully",
	})
}
