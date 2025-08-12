package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/controllers"
	"gofiberapp/middleware"
)

func StatusRoutes(router fiber.Router) {
	status := router.Group("/status")
	status.Use(middleware.SimpleAuth)
	status.Get("/", controllers.GetStatus)
	status.Get("/:id", controllers.GetStatuss)
	status.Post("/", controllers.CreateStatus)
	status.Patch("/:id", controllers.UpdateStatus)
	status.Delete("/:id", controllers.DeleteStatus)
}
