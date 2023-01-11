package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	db := database.DB
	new_user := new(models.User)

	if err := c.BodyParser(new_user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var user models.User

	db.Table("users").Where("email = ?", new_user.Email).First(&user)

	if user.ID != 0 {
		return c.Status(404).JSON(fiber.Map{
			"data":    nil,
			"message": "Email exist!",
		})
	}

	hash, _ := HashPassword(new_user.Password)
	new_user.Password = hash

	db.Create(&new_user)

	return c.Status(201).JSON(fiber.Map{
		"data":    new_user,
		"message": "Register account success!",
	})
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	// var data map[string]string
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db.Table("users").Where("email = ?", data["email"]).First(&user)

	log.Print("user", user)
	// log.Print("user", data["email"])

	return nil
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
