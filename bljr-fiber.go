package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello daffa")
	})

	fmt.Println("Fiber running on http://localhost:3000")
	app.Listen(3000)
}
