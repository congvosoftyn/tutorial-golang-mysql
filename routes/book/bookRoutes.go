package bookRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/handlers"
)

func SetupNoteRoutes(router fiber.Router) {
	book := router.Group("books")

	book.Get("/", handlers.GetAllBook)
	book.Post("/", handlers.CreateBook)
	book.Get("/:id", handlers.GetBook)
	book.Put("/:id", handlers.UpdateBook)
	book.Delete("/:id", handlers.DeleteBook)
}
