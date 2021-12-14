package beers

type Domain struct {
	ID       int64
	Name     string
	Brewery  string
	Country  string
	Price    float64
	Currency string
}

type CurrencyResponse struct {
	Success   bool
	Terms     string
	Privacy   string
	Timestamp int64
	Source    string
	Quotes    map[string]float64
}

type BeerBox struct {
	PriceTotal float64
}

type Service interface {
	Register(beer *Domain) (*Domain, error)
	Edit(beer *Domain, id int64) (*Domain, error)
	GetByID(id int64) (*Domain, error)
	All() (*[]Domain, error)
	BeerBox(beerID int64, quantity int, currency string) (*BeerBox, error)
}

type Repository interface {
	Create(beer *Domain) (*Domain, error)
	Update(beer *Domain, id int64) (*Domain, error)
	FindByID(id int64) (*Domain, error)
	All() (*[]Domain, error)
}
