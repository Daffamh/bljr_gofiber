package controllers

import (
	"github.com/gofiber/fiber"
	"gofiberapp/config"
	"gofiberapp/models"
)

func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(400).Send("gagal parsing data")
	}
	c.Send("Ganti dengan data dari DB ", id)
}

func CreateUser(c *fiber.Ctx) {
	user := new(models.User)

	err := c.BodyParser(&user)
	if err != nil {
		c.Status(400).Send(err.Error())
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.Status(400).Send(err.Error())
		return
	}
	config.DB.Create(&user)
	c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) {

	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(400).Send("gagal parsing data")
		return
	}
	c.Send("Ganti dengan data dari DB ", id)
}

func DeleteUser(c *fiber.Ctx) {
	id := c.Params("id")
	c.Send("DELETE: menghapus user dengan ID = " + id)
}
