package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/gofiber/fiber/v2"
)

func GetStudents(c *fiber.Ctx) error {
	var students []models.Student
	config.DB.Find(&students)
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student not found",
			"error":   err.Error(),
		})

	}
	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student

	err := c.BodyParser(&student)
	if err != nil {
		return c.Status(400).JSON(err.Error())

	}

	if err := config.DB.Create(&student).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}
	config.DB.Create(&student)
	return c.JSON(student)
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
	config.DB.Save(&student)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.Delete(&student, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete student",
			"error":   err.Error(),
		})

	}
	return c.JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}
