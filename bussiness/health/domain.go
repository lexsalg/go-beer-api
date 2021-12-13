package health

type Domain struct {
	Code   int
	Status string
}

type Service interface {
	Check() (*Domain, error)
}
