package main

import (
	"fmt"
	"gofiberapp/routes"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	fmt.Println("Fiber running on http://localhost:3000")
	app.Listen(3000)
}
