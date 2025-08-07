package controllers

import (
	"github.com/gofiber/fiber"
	"gofiberapp/config"
	"gofiberapp/models"
	"gofiberapp/utils"
	"golang.org/x/crypto/bcrypt"
)

// REGISTER
func Register(c *fiber.Ctx) {
	var input models.RegisterRequest
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		c.Status(400).JSON(err.Error())
		return
	}

	if input.Name == "" || input.Email == "" || input.Password == "" {
		err := c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Nama, Email dan Password wajib diisi",
		})
		if err != nil {
			c.Status(500).JSON(err.Error())
			return
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengenkripsi password",
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.PasswordHash = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mendaftarkan pengguna",
		})
		return
	}

	c.Status(200).JSON(input)
}

// LOGIN
func Login(c *fiber.Ctx) {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memproses data",
		})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Email atau password salah",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Email atau password salah",
		})
		return
	}

	// generate & update token
	token := utils.GenerateToken(32)

	config.DB.Model(&user).Update("token", token)

	c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"toen":   token,
		"data": fiber.Map{
			"id":    user.ID,
			"nama":  user.Name,
			"email": user.Email,
		},
	})
}
