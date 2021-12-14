package beers

import (
	"beerapi/app/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type serviceBeers struct {
	repository Repository
	currency   *config.CurrencyApi
}

func NewService(repoBeer Repository, currency *config.CurrencyApi) Service {
	return &serviceBeers{
		repository: repoBeer,
		currency:   currency,
	}
}

func (service serviceBeers) Register(beer *Domain) (*Domain, error) {
	result, err := service.repository.Create(beer)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}

func (service serviceBeers) Edit(beer *Domain, id int64) (*Domain, error) {

	result, err := service.repository.Update(beer, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}

func (service serviceBeers) All() (*[]Domain, error) {

	results, err := service.repository.All()
	if err != nil {
		list := []Domain{}
		return &list, err
	}
	return results, err
}

func (service serviceBeers) GetByID(id int64) (*Domain, error) {
	result, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}

func (service serviceBeers) BeerBox(beerId int64, quantity int, currency string) (*BeerBox, error) {

	beer, err := service.repository.FindByID(beerId)
	if err != nil {
		return &BeerBox{}, err
	}

	beerCurrency, err := service.requestApiCurrency(beer.Currency)

	if err != nil {
		return &BeerBox{}, err
	}

	beerCostDollar := (beer.Price / beerCurrency) * float64(quantity)

	// Para evitar error en el api currency: max limit request, plan free
	time.Sleep(2 * time.Second)

	payCurrency, err := service.requestApiCurrency(currency)

	if err != nil {
		return &BeerBox{}, err
	}

	beerCostCurrency := payCurrency * beerCostDollar

	return &BeerBox{PriceTotal: beerCostCurrency}, nil
}

func (service serviceBeers) requestApiCurrency(currency string) (float64, error) {

	srcCurrency := "USD"
	url := fmt.Sprintf("%s%s%s%s%s%s%s", service.currency.Url, "?access_key=", service.currency.ApiKey, "&source=", srcCurrency, "&currencies=", currency)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	TIMEOUT := 2 * time.Second
	client := &http.Client{Timeout: TIMEOUT}
	res, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	var data CurrencyResponse
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return 0, err
	}

	mapKey := fmt.Sprintf("%s%s", srcCurrency, strings.ToUpper(currency))

	return data.Quotes[mapKey], nil
}
