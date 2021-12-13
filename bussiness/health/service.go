package health

type serviceHealth struct {
}

func NewService() Service {
	return &serviceHealth{}
}

func (service serviceHealth) Check() (*Domain, error) {

	return &Domain{Code: 200, Status: "Ok"}, nil
}
