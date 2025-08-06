package routes

import (
	"github.com/gofiber/fiber"
	"gofiberapp/controllers"
)

func StudentRoutes(app *fiber.App) {
	student := app.Group("/student")
	student.Get("/", controllers.GetStudents)
	student.Get("/:id", controllers.GetStudent)
	student.Post("/", controllers.CreateStudent)
	student.Patch("/:id", controllers.UpdateStudent)
	student.Delete("/:id", controllers.DeleteStudent)
}
