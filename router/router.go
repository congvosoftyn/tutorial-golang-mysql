package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vothecong/go-tutorial/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())
	book := api.Group("books")

	book.Get("/", handlers.GetAllBook)
	book.Post("/", handlers.CreateBook)
	book.Get("/:id", handlers.GetBook)
	book.Put("/:id", handlers.UpdateBook)
	book.Delete("/:id", handlers.DeleteBook)
}
