package main

import (
	"gofiberapp/config"
	"gofiberapp/models"
	"gofiberapp/routes"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.Student{}, &models.User{})
	if err != nil {
		println(err.Error())
		return
	}

	routes.Routes(app)

	println("Fiber running on http://localhost:3000")
	err = app.Listen(3000)
	if err != nil {
		println(err.Error())
		return
	}
}
