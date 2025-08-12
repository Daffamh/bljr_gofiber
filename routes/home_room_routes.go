package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiberapp/controllers"
	"gofiberapp/middleware"
)

func HomeRoomRoutes(router fiber.Router) {
	homeroom := router.Group("/home_room")
	homeroom.Use(middleware.SimpleAuth)
	homeroom.Get("/", controllers.GetHomeRoom)
	homeroom.Get("/:id", controllers.GetHomeRooms)
	homeroom.Post("/", controllers.CreateHomeRoom)
	homeroom.Patch("/:id", controllers.UpdateHomeRoom)
	homeroom.Delete("/:id", controllers.DeleteHomeRoom)
}
