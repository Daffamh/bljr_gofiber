package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/controllers"
	"gofiberapp/middleware"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
	auth.Use(middleware.SimpleAuth)
	auth.Get("/logout", controllers.Logout)
	auth.Get("/me", controllers.GetProfile)
}
