package main

import (
	"github.com/gofiber/fiber/v2"

	"gofiberapp/config"
	"gofiberapp/models"
	"gofiberapp/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.StudentGrade{}, &models.Status{}, &models.HomeRoom{}, &models.Grade{}, &models.Student{}, &models.User{})
	if err != nil {
		println(err.Error())
		return
	}

	routes.Routes(app)

	println("Fiber running on http://localhost:3000")
	err = app.Listen(":3000")
	if err != nil {
		println(err.Error())
		return
	}
}
