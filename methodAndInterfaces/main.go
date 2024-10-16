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

type Account struct {
	accountNumber int
	expences      []Expencive
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

func (product Product) getName() string {
	return product.name
}

func (product Product) getCoast(resurse bool) float64 {
	return product.price
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
func main() {
	account := Account{
		accountNumber: 25552,
		expences: []Expencive{
			Product{"Kayak", "Watersport", 275},
			Service{"Bebra Corp", 12, 87.25},
		},
	}
	fmt.Println(account.accountNumber)
	for _, expence := range account.expences {
		fmt.Println(expence.getName(), expence.getCoast(true))
	}
	fmt.Println(calcTotal(account.expences))
}
