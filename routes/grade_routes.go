package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/controllers"
	"gofiberapp/middleware"
)

func GradeRoutes(router fiber.Router) {
	grade := router.Group("/grade")
	grade.Use(middleware.SimpleAuth)
	grade.Get("/", controllers.GetGrade)
	grade.Get("/:id", controllers.GetGrades)
	grade.Post("/", controllers.CreateGrade)
	grade.Patch("/:id", controllers.UpdateGrade)
	grade.Delete("/:id", controllers.DeleteGrade)
}
