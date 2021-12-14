package beers

import (
	"beerapi/app/presenter/beers/request"
	"beerapi/app/presenter/beers/response"
	"beerapi/bussiness/beers"
	"beerapi/helper/response_tmp"
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Presenter struct {
	serviceBeer beers.Service
}

func NewHandler(beerService beers.Service) *Presenter {
	return &Presenter{
		serviceBeer: beerService,
	}
}

func (handler *Presenter) Register(fiberContext *fiber.Ctx) error {
	var req request.Beer
	if err := fiberContext.BodyParser(&req); err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusBadRequest, "error", err)
		return fiberContext.Status(http.StatusBadRequest).JSON(response)
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceBeer.Register(domain)
	if err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success Register Beer", http.StatusOK, "Success", response.FromDomain(*resp))
	return fiberContext.JSON(responseRes)
}

func (handler *Presenter) Update(fiberContext *fiber.Ctx) error {
	beerID, err := fiberContext.ParamsInt("beerID")
	if err != nil {
		response := response_tmp.APIResponse("Beer Id not valid.", http.StatusBadRequest, "bad request", err)
		return fiberContext.JSON(response)
	}

	var req request.Beer
	if err := fiberContext.BodyParser(&req); err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusBadRequest, "error", err)
		return fiberContext.Status(http.StatusBadRequest).JSON(response)
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceBeer.Edit(domain, int64(beerID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := response_tmp.APIResponse("Beer Not found", http.StatusNotFound, "not found", err)
			return fiberContext.Status(http.StatusNotFound).JSON(response)
		}
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success Register Beer", http.StatusOK, "Success", response.FromDomain(*resp))
	return fiberContext.JSON(responseRes)
}

func (handler *Presenter) All(fiberContext *fiber.Ctx) error {

	resp, err := handler.serviceBeer.All()
	if err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success", http.StatusOK, "Success", response.FromDomainList(*resp))
	return fiberContext.JSON(responseRes)
}

func (handler *Presenter) GetById(fiberContext *fiber.Ctx) error {
	beerID, err := fiberContext.ParamsInt("beerID")
	if err != nil {
		response := response_tmp.APIResponse("Beer Id not valid.", http.StatusBadRequest, "bad request", err)
		return fiberContext.JSON(response)
	}

	resp, err := handler.serviceBeer.GetByID(int64(beerID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := response_tmp.APIResponse("Beer Not found", http.StatusNotFound, "not found", err)
			return fiberContext.Status(http.StatusNotFound).JSON(response)
		}
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success", http.StatusOK, "Success", response.FromDomain(*resp))
	return fiberContext.JSON(responseRes)
}

func (handler *Presenter) BeerBox(fiberContext *fiber.Ctx) error {
	beerID, err := fiberContext.ParamsInt("beerID")
	if err != nil {
		response := response_tmp.APIResponse("Beer Id not valid.", http.StatusBadRequest, "bad request", err)
		return fiberContext.JSON(response)
	}

	quantityStr := fiberContext.Query("quantity")
	quantity := 6
	if quantityStr != "" {
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			response := response_tmp.APIResponse("Quantity not valid.", http.StatusBadRequest, "bad request", err)
			return fiberContext.JSON(response)
		}
	}
	currency := fiberContext.Query("currency")
	if currency == "" {
		response := response_tmp.APIResponse("Currency not valid.", http.StatusBadRequest, "bad request", "")
		return fiberContext.JSON(response)
	}

	resp, err := handler.serviceBeer.BeerBox(int64(beerID), quantity, currency)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := response_tmp.APIResponse("Beer Not found", http.StatusNotFound, "not found", err)
			return fiberContext.Status(http.StatusNotFound).JSON(response)
		}
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success", http.StatusOK, "Success", response.BeerBox{PriceTotal: resp.PriceTotal})
	return fiberContext.JSON(responseRes)
}
