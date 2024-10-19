package main

import (
	"fmt"
)

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

// type Expencive interface {
// 	getName() string
// 	getCoast(annual bool) float64
// }

// func (p Product) getName() string {
// 	return p.name
// }

// func (p Product) getCoast(_ bool) float64 {
// 	return p.price
// }

// func (s Service) getName() string {
// 	return s.description
// }

// func (s Service) getCoast(recur bool) float64 {
// 	if recur {
// 		return float64(s.durrationMonths) * s.monthlyFee
// 	}
// 	return s.monthlyFee
// }

// func main() {
// 	expences := []Expencive{
// 		Product{"Kayak", "Watersport", 275},
// 		Service{"Boat Cover", 12, 89.52},
// 	}
// 	for _, expence := range expences {
// 		fmt.Println(expence.getName(), expence.getCoast(true))
// 	}
// }

type Expencive interface {
	getName() string
	getCoast(annual bool) float64
}

type Person struct {
	name, city string
}

type Account struct {
	accountNumber int
	expences      []Expencive
}

func calcTotal(expences []Expencive) (total float64) {
	for _, item := range expences {
		total += item.getCoast(true)
	}
	return total
}

//	func main() {
//		expences := []Expencive{
//			Product{"Kayak", "Watersport", 275},
//			Service{"Boyler Corp", 12, 87.52},
//		}
//		for _, expence := range expences {
//			fmt.Println(expence.getName(), expence.getCoast(true))
//		}
//		fmt.Println("Total price", calcTotal(expences))
//	}
//
//	func main() {
//		account := Account{
//			accountNumber: 25552,
//			expences: []Expencive{
//				Product{"Kayak", "Watersport", 275},
//				Service{"Bebra Corp", 12, 87.25},
//			},
//		}
//		fmt.Println(account.accountNumber)
//		for _, expence := range account.expences {
//			fmt.Println(expence.getName(), expence.getCoast(true))
//		}
//		fmt.Println(calcTotal(account.expences))
//	}
//
//	func main() {
//		product := Product{"Kayak", "Watersport", 275}
//		var expence Expencive = &product
//		product.price = 100
//		fmt.Println(product.price)
//		fmt.Println(expence.getCoast(true))
//	}
// func main() {
// 	var e1 Expencive = &Product{name: "Kayak"}
// 	var e2 Expencive = &Product{name: "Kayak"}
// 	//e5 := &e1
// 	var e3 Expencive = Service{description: "Bebra Corp"}
// 	var e4 Expencive = Service{description: "Bebra Corp"}
// 	fmt.Println("e1 == e2", e1 == e2) //сравниваются места в памяти
// 	fmt.Println("e3 == e4?", e3 == e4)
// 	//fmt.Println("e1 == e5?", e1 == *e5)

// }

// func main() {
// 	expenses := []Expencive{
// 		Service{"Bebra Corp", 12, 89.50, []string{}},
// 		Service{"Padal Protact", 12, 8, []string{}},
// 		&Product{"Kayak", "Watersport", 285},
// 	}
// 	// for _, expense := range expenses {
// 	// 	if s, ok := expense.(Service); ok {
// 	// 		fmt.Println(s.description, s.monthlyFee*float64(s.durrationMonths))
// 	// 	} else {
// 	// 		fmt.Println(expense.getName(), expense.getCoast(true))
// 	// 	}
// 	// }
// 	for _, expense := range expenses {
// 		switch value := expense.(type) {
// 		case Service:
// 			fmt.Println("Service:", value.description, value.monthlyFee*float64(value.durrationMonths))
// 		case *Product:
// 			fmt.Println("Product:", value.name, value.price)
// 		default:
// 			fmt.Println(expense.getName(), expense.getCoast(true))
// 		}
// 	}
// }

func processItem(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, value.price)
		case *Product:
			fmt.Println("Product pointer:", value.name, value.price)
		case Service:
			fmt.Println("Service:", value.description, value.monthlyFee*float64(value.durrationMonths))
		case string:
			fmt.Println(value)
		case int:
			fmt.Println("This is int value:", value)
		case bool:
			fmt.Println("This is bool value", value)
		case Person:
			fmt.Println("Person:", value.name, value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, value.city)
		default:
			fmt.Println(item)
		}
	}
}

func main() {
	var expense Expencive = &Product{"Kayak", "Watersport", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersport", 45.52},
		Service{"Bebra Corp", 12, 89.5, []string{}},
		Person{"Alice", "Volgodonsk"},
		&Person{"Bober", "Warshava"},
		"This is a string",
		100,
		true,
	}
	//for _, item := range data {
	//processItem(item)
	// switch value := item.(type) {
	// case Product:
	// 	fmt.Println("Product:", value.name, value.price)
	// case *Product:
	// 	fmt.Println("Product pointer:", value.name, value.price)
	// case Service:
	// 	fmt.Println("Service:", value.description, value.monthlyFee*float64(value.durrationMonths))
	// case string:
	// 	fmt.Println(value)
	// case int:
	// 	fmt.Println("This is int value:", value)
	// case bool:
	// 	fmt.Println("This is bool value", value)
	// case Person:
	// 	fmt.Println("Person:", value.name, value.city)
	// case *Person:
	// 	fmt.Println("Person Pointer:", value.name, value.city)
	// default:
	// 	fmt.Println(expense.getName(), expense.getCoast(true))
	// }
	processItem(data...)
}
