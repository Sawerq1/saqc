package main

import "fmt"

// type Product struct {
// 	name, category string
// 	price          float64
// }

// type Suppler struct {
// 	name, city string
// }

// func (suppler *Suppler) printDetails() {
// 	fmt.Println(suppler.name, suppler.city)
// }

// func newSuppler(name, city string) *Suppler {
// 	return &Suppler{name, city}
// }

// func printDetails(product *Product) {
// 	fmt.Println(product.name, product.category, product.price)
// }

// func newProduct(name, category string, price float64) *Product {
// 	return &Product{name, category, price}
// }

// func (product *Product) printDetails() {
// 	fmt.Println(product.name, product.category, product.price, product.calcTax(0.2, 100))
// }

// func (product Product) printDetails() {
// 	fmt.Println(product.name, product.category, product.price, product.calcTax(0.2, 100))
// }

// func (product *Product) calcTax(rate, threshold float64) float64 {
// 	if product.price > threshold {
// 		return product.price*rate + product.price
// 	}
// 	return product.price
// }

//	func main() {
//		kayak := &Product{"Kayak", "Watersport", 275}
//		kayak.printDetails()
//	}

func main() {
	kayak := Product{"Kayak", "Watersport", 275}
	insurence := Service{"Boat Cover", 12, 89.52}
	fmt.Println("Product:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
	fmt.Println("Service:", insurence.description, "Price:", insurence.monthlyFee*float64(insurence.durrationMonths))
}
