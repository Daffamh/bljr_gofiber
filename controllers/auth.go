package controllers

import (
	"gofiberapp/config"
	"gofiberapp/models"
	"gofiberapp/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// REGISTER
func Register(c *fiber.Ctx) error {
	var input models.RegisterRequest
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if input.Name == "" || input.Email == "" || input.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Nama, Email dan Password wajib diisi",
		})

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengenkripsi password",
			"error":   err.Error(),
		})

	}

	user.Name = input.Name
	user.Email = input.Email
	user.PasswordHash = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mendaftarkan pengguna",
			"error":   err.Error(),
		})

	}

	return c.Status(200).JSON(input)
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Gagal memproses data",
		})

	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Email atau password salah",
			"error":   err.Error(),
		})

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Email atau password salah",
			"error":   err.Error(),
		})

	}

	// generate & update token
	token := utils.GenerateToken(32)

	config.DB.Model(&user).Update("token", token)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"token":   token,
		"data": fiber.Map{
			"id":    user.ID,
			"nama":  user.Name,
			"email": user.Email,
		},
	})
}

func Logout(c *fiber.Ctx) error {

	id := c.Locals("id")
	config.DB.Model(&models.User{}).Where("id = ?", id).Update("token", "")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Logout berhasil",
	})
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("id").(uint) // Get user ID from middleware

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
