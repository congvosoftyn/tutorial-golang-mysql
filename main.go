package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/database"
	"github.com/vothecong/go-tutorial/router"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	database.ConnectionDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
