package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/")
	UserRoutes(api)
	StudentRoutes(api)
	AuthRoutes(api)
	GradeRoutes(api)
	HomeRoomRoutes(api)
	StatusRoutes(api)
}
