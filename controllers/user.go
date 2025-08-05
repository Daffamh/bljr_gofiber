package controllers

import (
	"github.com/gofiber/fiber"
	"gofiberapp/models"
)

func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(400).Send("gagal parsing data")
	}
	c.Send("GET:menampilkan user dengan nama = " + user.Nama + ", alamat = " + user.Alamat + ", ID = " + id)
}

func CreateUser(c *fiber.Ctx) {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(400).Send("gagal presing data")
	}
	c.Send("POST:menambahkan user = " + user.Nama + ", alamat = " + user.Alamat)
}

func UpdateUser(c *fiber.Ctx) {

	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(400).Send("gagal parsing data")
		return
	}
	c.Send("PATCH: update user ID = " + id + ", nama = " + user.Nama + ", alamat = " + user.Alamat)
}

func DeleteUser(c *fiber.Ctx) {
	id := c.Params("id")
	c.Send("DELETE: menghapus user dengan ID = " + id)
}
