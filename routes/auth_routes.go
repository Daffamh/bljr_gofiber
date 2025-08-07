package routes

import (
	"gofiberapp/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
