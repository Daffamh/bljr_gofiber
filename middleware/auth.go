package middleware

import (
	"gofiberapp/config"
	"gofiberapp/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SimpleAuth(c *fiber.Ctx) {
	email := c.Get("X_User_email")
	password := c.Get("X_User_Password")

	if email == "" || password == "" {
		c.Status(401).SendString("Authentical diperlukan")
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.Status(401).SendString("User tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.Status(401).SendString("Password salah")
		return
	}
	c.Locals("user", user)
	c.Next()
}
