package main

type Product struct {
	name, category string
	price          float64
}

func (product *Product) getName() string {
	return product.name
}

func (product *Product) getCoast(resurse bool) float64 {
	return product.price
}
