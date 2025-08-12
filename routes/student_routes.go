package routes

import (
	"gofiberapp/controllers"
	"gofiberapp/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(router fiber.Router) {
	student := router.Group("/student")
	student.Use(middleware.SimpleAuth)
	student.Get("/", controllers.GetStudents)
	student.Get("/:id", controllers.GetStudent)
	student.Post("/", controllers.CreateStudent)
	student.Patch("/:id", controllers.UpdateStudent)
	student.Delete("/:id", controllers.DeleteStudent)
}
