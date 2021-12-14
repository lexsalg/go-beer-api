package request

import "beerapi/bussiness/beers"

type Beer struct {
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency" validate:"required,min=3,max=3"`
}

func ToDomain(request Beer) *beers.Domain {
	return &beers.Domain{
		Name:     request.Name,
		Brewery:  request.Brewery,
		Country:  request.Country,
		Price:    request.Price,
		Currency: request.Currency,
	}
}
