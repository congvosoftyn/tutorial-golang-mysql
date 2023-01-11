package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	bookRoutes "github.com/vothecong/go-tutorial/routes/book"
	userRoutes "github.com/vothecong/go-tutorial/routes/user"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())

	bookRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
}
