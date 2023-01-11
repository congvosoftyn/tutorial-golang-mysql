package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vothecong/go-tutorial/handlers"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("users")

	user.Post("/register", handlers.Register)
	user.Post("/login", handlers.Login)
	user.Post("/logout", handlers.Logout)

}
