package main

import "fmt"

// func main() {
// 	type Product struct {
// 		name, category string
// 		price          float64
// 	}

// 	type StockLevel struct {
// 		Product
// 		Alternative Product
// 		Alt         Product
// 		count       int
// 	}

//	stockItem := StockLevel{
//		Product:     Product{"Kayak", "Watersports", 275},
//		Alternative: Product{"Lifejacket", "Watersports", 48.95},
//		Alt:         Product{"Soccer Ball", "Football", 5.5},
//		count:       100,
//	}
//	fmt.Println(stockItem.Product.name, stockItem.count)
//	fmt.Println(stockItem.Alternative.name, stockItem.Alternative.price)
//	fmt.Println(stockItem.Alt.name, stockItem.Alt.price)
//
//	kayak := Product{
//		name:     "Kayak",
//		category: "Watersports",
//		//price:    275,
//	}
//
// kayak := Product{"Kayak", "Watersports", 275}
// fmt.Println(kayak.name, kayak.category, kayak.price)
// kayak.price = 300
// fmt.Println("Changed price:", kayak.price)
// var lifejacket Product
// fmt.Println(lifejacket.name, lifejacket.category, lifejacket.price)
// }
// func writeName(val struct {
// 	name, category string
// 	price          float64
// }) {
// 	fmt.Println(val.name)
// }

// func main() {
// 	type Product struct {
// 		name, category string
// 		price          float64
// 		//overNames		[]string
// 	}

// 	type Item struct {
// 		name     string
// 		category string
// 		price    float64
// 	}

// 	prod := Product{name: "Kayak", category: "Watersport", price: 275}
// 	item := Item{name: "Stadium", category: "Sacker", price: 750000}
// 	writeName(prod)
// 	writeName(item)
//fmt.Println("p1 == p2?", prod == Product(item))

//}

type Product struct {
	name, category string
	price          float64
	*Suppler
}

type Suppler struct {
	name, sity string
}

func newProduct(name, category string, price float64, suppler *Suppler) *Product {
	return &Product{name, category, price, suppler}
}

func copyProduct(product *Product) Product {
	p := *product
	s := *product.Suppler
	p.Suppler = &s
	return p
}

func main() {
	acme := &Suppler{"Acme", "New-York"}
	p1 := newProduct("Kayak", "Watersport", 275, acme)
	p2 := copyProduct(p1)
	p1.name = "Original Kayak"
	p1.Suppler.name = "Humster"
	for _, p := range []Product{*p1, p2} {
		fmt.Println(p.name, p.Suppler.sity, p.Suppler.name)
	}
}
