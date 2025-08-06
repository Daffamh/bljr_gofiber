package controllers

import (
	"github.com/gofiber/fiber"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetStudents(c *fiber.Ctx) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(students)
}

func GetStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.Status(404).Send("student tidak ditemukan")
		return
	}
	c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) {
	var student models.Student

	err := c.BodyParser(&student)
	if err != nil {
		c.Status(400).Send(err.Error())
		return
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.Status(400).Send(err.Error())
		return
	}
	config.DB.Create(&student)
	c.JSON(student)
}

func UpdateStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.Status(404).Send("student tidak ditemukan")
		return
	}
	if err := c.BodyParser(&student); err != nil {
		c.Status(400).Send("data tidak valid")
		return
	}
	config.DB.Save(&student)
	c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	if err := config.DB.Delete(&student, id).Error; err != nil {
		c.Status(500).Send("Gagal hapus student")
		return
	}
	c.Send("Berhasil dihapus")
}
