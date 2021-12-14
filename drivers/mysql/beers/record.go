package beers

import (
	"beerapi/bussiness/beers"

	"gorm.io/gorm"
)

type Beers struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string
	Brewery  string
	Country  string
	Price    float64
	Currency string
}

func toDomain(record Beers) beers.Domain {
	return beers.Domain{
		ID:       int64(record.ID),
		Name:     record.Name,
		Brewery:  record.Brewery,
		Country:  record.Country,
		Price:    record.Price,
		Currency: record.Currency,
	}
}

func toDomainList(records []Beers) []beers.Domain {
	list := []beers.Domain{}
	for _, item := range records {
		domain := toDomain(item)
		list = append(list, domain)
	}
	return list
}

func fromDomain(domain beers.Domain) Beers {
	return Beers{
		ID:       uint64(domain.ID),
		Name:     domain.Name,
		Brewery:  domain.Brewery,
		Country:  domain.Country,
		Price:    domain.Price,
		Currency: domain.Currency,
	}
}
