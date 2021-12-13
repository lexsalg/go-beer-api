package users

import (
	"beerapi/app/presenter/users/request"
	"beerapi/app/presenter/users/response"
	"beerapi/bussiness/users"
	"beerapi/helper/response_tmp"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Presenter struct {
	serviceUser users.Service
}

func NewHandler(userService users.Service) *Presenter {
	return &Presenter{
		serviceUser: userService,
	}
}

func (handler *Presenter) Register(fiberContext *fiber.Ctx) error {
	var req request.User
	if err := fiberContext.BodyParser(&req); err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusBadRequest, "error", err)
		return fiberContext.JSON(response)
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Register(domain)
	if err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success Register User", http.StatusOK, "Success", response.FromDomain(*resp))
	return fiberContext.JSON(responseRes)
}
