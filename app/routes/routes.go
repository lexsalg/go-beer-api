package routes

import (
	"beerapi/app/presenter/beers"
	"beerapi/app/presenter/health"
	"beerapi/app/presenter/users"

	"github.com/gofiber/fiber/v2"
)

type HandlerList struct {
	HealthHandler health.Presenter
	UserHandler   users.Presenter
	BeerHandler   beers.Presenter
}

func (handler *HandlerList) Routes(fiberContext *fiber.App) {
	api := fiberContext.Group("/api/v1")

	//health endpoint
	api.Get("/health", handler.HealthHandler.Check)

	//user Endpoint
	api.Post("/users/register", handler.UserHandler.Register)

	//beers Endpoint
	beers := api.Group("/beers")
	beers.Get("", handler.BeerHandler.All)
	beers.Post("", handler.BeerHandler.Register)
	beers.Put("/:beerID", handler.BeerHandler.Update)
	beers.Get("/:beerID", handler.BeerHandler.GetById)
	beers.Get("/:beerID/boxprice", handler.BeerHandler.BeerBox)

}
