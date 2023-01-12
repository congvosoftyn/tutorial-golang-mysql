package handlers

import (
	"errors"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vothecong/go-tutorial/config"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getUserByEmail(email string) (*models.User, error) {
	db := database.DB
	var user models.User

	if err := db.Where(&models.User{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &user, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Register(c *fiber.Ctx) error {
	db := database.DB
	new_user := new(models.User)

	if err := c.BodyParser(new_user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// db.Table("users").Where("email = ?", new_user.Email).First(&user)
	email, err := getUserByEmail(new_user.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if email.ID != 0 {
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

	match := CheckPassword(data["password"], user.Password)

	if !match {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"message": "Email or Password incorect!",
			"error":   fiber.StatusBadRequest,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	_token, err := token.SignedString([]byte(config.Config("SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"data":    _token,
		"message": nil,
		"error":   nil,
	})
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
