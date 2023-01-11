package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/models"
)

func GetAllBook(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Select([]string{"title", "description"}).Find(&books)
	return c.Status(200).JSON(fiber.Map{"data": books})
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Save(&book)

	return c.Status(201).JSON(fiber.Map{
		"data":   book,
		"status": "success",
	})
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	// result := database.DB.Find(&book, id)
	// result := database.DB.First(&book, "id = ?", id)
	result := database.DB.First(&book).Where("id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":  nil,
			"error": "Not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": book,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book = new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&book)

	return c.Status(200).JSON(fiber.Map{
		"data": book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := database.DB.Delete(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
