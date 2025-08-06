package routes

import (
	"github.com/gofiber/fiber"
)

func Routes(app *fiber.App) {
	UserRoutes(app)
	StudentRoutes(app)
}
