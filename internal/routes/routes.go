package routes

import (
	"github.com/Anagh3/go-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, handler *handlers.UserHandler) {
	app.Post("/users", handler.CreateUserHandler)
	app.Get("/users", handler.GetAllUsersHandler)
	app.Get("/users/:id", handler.GetUserByIDHandler)
	app.Put("/users/:id", handler.UpdateUserHandler)
	app.Delete("/users/:id", handler.DeleteUserHandler)
}
