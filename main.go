package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/models"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	database.ConnectionDB()

	app.Get("/books", func(c *fiber.Ctx) error {
		var books []models.Book
		database.DB.Find(&books)
		return c.Status(200).JSON(fiber.Map{"data": books})
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		book := new(models.Book)

		if err := c.BodyParser(book); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		database.DB.Save(&book)

		return c.Status(201).JSON(fiber.Map{
			"data":   book,
			"status": "success",
		})
	})

	app.Get("/book/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.Book

		result := database.DB.Find(&book, id)

		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"data":  nil,
				"error": "Not found",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"data": book,
		})
	})

	app.Put("/book/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book = new(models.Book)

		if err := c.BodyParser(book); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		database.DB.Where("id = ?", id).Updates(&book)

		return c.Status(200).JSON(fiber.Map{
			"data": book,
		})
	})
	app.Delete("/book/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.Book

		result := database.DB.Delete(&book, id)

		if result.RowsAffected == 0 {
			return c.SendStatus(404)
		}

		return c.SendStatus(200)
	})
	app.Listen(":3000")
}
