package routes

import (
	"beerapi/app/presenter/health"
	"beerapi/app/presenter/users"

	"github.com/gofiber/fiber/v2"
)

type HandlerList struct {
	UserHandler   users.Presenter
	HealthHandler health.Presenter
}

func (handler *HandlerList) Routes(fiberContext *fiber.App) {
	api := fiberContext.Group("/api/v1")

	//health endpoint
	api.Get("/health", handler.HealthHandler.Check)

	//user Endpoint
	api.Post("/users/register", handler.UserHandler.Register)
}
