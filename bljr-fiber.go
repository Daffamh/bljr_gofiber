package main

import (
	"fmt"
	"gofiberapp/config"
	"gofiberapp/models"
	"gofiberapp/routes"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Student{})

	routes.Routes(app)

	fmt.Println("Fiber running on http://localhost:3000")
	app.Listen(3000)
}
