package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetStudents(c *fiber.Ctx) error {
	var students []models.Student
	config.DB.Unscoped().
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		Find(&students)
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		//Preload("StudentGrade.HomeRoom").
		//Preload("StudentGrade.Grade").
		First(&student, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student
	validate := validator.New()

	err := c.BodyParser(&student)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid student",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(student)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("id").(uint)
	student.CreatedById = &userID
	student.UpdatedById = &userID

	if err := config.DB.Create(&student).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStudent models.Student
	if err := config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		First(&createdStudent, student.Id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(createdStudent)
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student not found",
			"error":   err.Error(),
		})

	}
	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})

	}

	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("id").(uint)
	student.UpdatedById = &userID

	config.DB.Save(&student)

	var updatedStudent models.Student
	config.DB.
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("DeletedBy").
		Preload("StudentGrade").
		Where("id = ?", id).First(&updatedStudent)

	return c.JSON(updatedStudent)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	if err := config.DB.Delete(&models.Student{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete student",
			"error":   err.Error(),
		})
	}

	if err := config.DB.Unscoped().Model(&models.Student{}).Where("id = ?", id).Update("deleted_by_id", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update deleted_by field",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Student deleted successfully",
	})
}
