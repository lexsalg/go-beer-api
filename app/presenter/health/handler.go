package health

import (
	"beerapi/bussiness/health"

	"beerapi/helper/response_tmp"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Presenter struct {
	serviceHealth health.Service
}

func NewHandler(healthService health.Service) *Presenter {
	return &Presenter{
		serviceHealth: healthService,
	}
}

func (handler *Presenter) Check(fiberContext *fiber.Ctx) error {

	res, err := handler.serviceHealth.Check()
	if err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}

	return fiberContext.JSON(res)
}
