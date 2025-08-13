package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetStudents(c *fiber.Ctx) error {
	var students []models.Student
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		Find(&students)
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
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
	student.CreatedBy = userID
	student.UpdatedBy = userID

	if err := config.DB.Create(&student).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data request",
			"error":   err.Error(),
		})
	}

	var createdStudent models.Student
	if err := config.DB.
		Preload("Creator").
		Preload("Updater").
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
	student.UpdatedBy = userID

	config.DB.Save(&student)

	var updatedStudent models.Student
	config.DB.
		Preload("Creator").
		Preload("Updater").
		Preload("StudentGrade").
		Where("id = ?", id).First(&updatedStudent)

	return c.JSON(updatedStudent)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("id").(uint)

	var student models.Student
	student.DeletedBy = &userID
	if err := config.DB.Delete(&student, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete student",
			"error":   err.Error(),
		})

	}

	if err := config.DB.Save(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete student",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}
