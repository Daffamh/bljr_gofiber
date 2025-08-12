package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/controllers"
	"gofiberapp/middleware"
)

func StudentGradeRoutes(router fiber.Router) {
	studentGrade := router.Group("/student_grade")
	studentGrade.Use(middleware.SimpleAuth)
	studentGrade.Get("/", controllers.GetStudentGrade)
	studentGrade.Get("/:id", controllers.GetStudentGrades)
	studentGrade.Post("/", controllers.CreateStudentGrade)
	studentGrade.Patch("/:id", controllers.UpdateStudentGrade)
	studentGrade.Delete("/:id", controllers.DeleteStudentGrade)
}
