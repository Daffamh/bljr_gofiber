package routes

import (
	"gofiberapp/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
	user.Get("/:id/students", controllers.GetUserWithStudents)
}
