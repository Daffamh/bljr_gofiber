package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

type user struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

func main() {
	app := fiber.New()

	app.Get("/user/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		user := new(user)

		if err := c.BodyParser(user); err != nil {
			c.Status(400).Send("gagal parsing data")
		}
		c.Send("GET:menampilkan user dengan nama = " + user.Nama + ", alamat = " + user.Alamat + ", ID = " + id)
	})

	app.Post("/user", func(c *fiber.Ctx) {
		user := new(user)

		if err := c.BodyParser(user); err != nil {
			c.Status(400).Send("gagal presing data")
		}
		c.Send("POST:menambahkan user = " + user.Nama + ", alamat = " + user.Alamat)
	})

	app.Put("user/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		user := new(user)

		if err := c.BodyParser(user); err != nil {
			c.Status(400).Send("gagal parsing data")
		}
		c.Send("PUT: update user ID = " + id + ", nama = " + user.Nama + ", alamat = " + user.Alamat)
	})

	app.Patch("user/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		user := new(user)
		if err := c.BodyParser(user); err != nil {
			c.Status(400).Send("gagal parsing data")
			return
		}
		c.Send("PATCH: update user ID = " + id + ", nama = " + user.Nama + ", alamat = " + user.Alamat)
	})

	app.Delete("user/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		c.Send("DELETE: menghapus user dengan ID = " + id)
	})

	fmt.Println("Fiber running on http://localhost:3000")
	app.Listen(3000)
}
