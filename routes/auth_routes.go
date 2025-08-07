package routes

import (
	"github.com/gofiber/fiber"
	"gofiberapp/controllers"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
