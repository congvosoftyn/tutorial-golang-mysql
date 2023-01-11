package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	database.ConnectionDB()

	app.Get("/books", handlers.GetAllBook)
	app.Post("/books", handlers.CreateBook)
	app.Get("/book/:id", handlers.GetBook)
	app.Put("/book/:id", handlers.UpdateBook)
	app.Delete("/book/:id", handlers.DeleteBook)

	app.Listen(":3000")
}
