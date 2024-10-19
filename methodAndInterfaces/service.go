package main

type Service struct {
	description     string
	durrationMonths int
	monthlyFee      float64
	featyres        []string
}

func (service Service) getName() string {
	return service.description
}

func (service Service) getCoast(res bool) float64 {
	if res {
		return service.monthlyFee * float64(service.durrationMonths)
	}
	return service.monthlyFee
}
