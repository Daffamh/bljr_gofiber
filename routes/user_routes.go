package routes

import (
	"github.com/gofiber/fiber"
	"gofiberapp/controllers"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}
