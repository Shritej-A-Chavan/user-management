package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-task/internal/handler"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Get("/users", userHandler.GetUsers)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
