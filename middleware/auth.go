package middleware

import (
	"gofiberapp/config"
	"gofiberapp/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SimpleAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	token := strings.Split(header, "Bearer ")[1]

	var user models.User
	if err := config.DB.Where("token = ?", token).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "User tidak ditemukan",
			"error":   err.Error(),
		},
		)
	}

	c.Locals("id", user.ID)
	return c.Next()
}
