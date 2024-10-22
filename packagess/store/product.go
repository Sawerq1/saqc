/*
Package store have types and methods
for online salles
*/

package store

// Product discribe an item for sale
type Product struct {
	Name, Category string // name and type of product
	price          float64
}

var standartTax = newTaxRate(0.25, 50)

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return standartTax.calcTax(p)
}

func (p *Product) SetPrice(NewPrice float64) {
	p.price = NewPrice
}
