package response

import (
	"beerapi/bussiness/beers"
)

type Beer struct {
	ID       int64   `json:"Id"`
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency"`
}

type BeerBox struct {
	PriceTotal float64 `json:"Price Total"`
}

func FromDomain(domain beers.Domain) Beer {
	return Beer{
		ID:       domain.ID,
		Name:     domain.Name,
		Brewery:  domain.Brewery,
		Country:  domain.Country,
		Price:    domain.Price,
		Currency: domain.Currency,
	}
}

func FromDomainList(records []beers.Domain) []Beer {
	list := []Beer{}
	for _, item := range records {
		domain := FromDomain(item)
		list = append(list, domain)
	}
	return list
}
