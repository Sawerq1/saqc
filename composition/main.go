package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// kayak := store.NewProduct("Kayak", "Watersport", 275)
	// // cdfmt.Println(kayak.Name, kayak.Category, kayak.Price(0.2))
	// lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersport"}
	// for _, p := range []*store.Product{kayak, lifejacket} {
	// 	fmt.Println("name:", p.Name, "category:", p.Category, "price with tax:", p.Price(0.2))
	// }
	// boats := []*store.Boat{
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Tender", 650.25, 2, true),
	// }
	// for _, boat := range boats {
	// 	fmt.Println(boat.Name, boat.Capacity, boat.Motorized, boat.Price(0.2))
	// }
	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Yach", 50000, 5, true, true, "Bobby", "Charle"),
	// 	store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
	// 	store.NewRentalBoat("Super Yach", 100000, 15, true, true, "Alice", "Shlapa"),
	// }
	// for _, r := range rentals {
	// 	if r.IncludeCrew {
	// 		fmt.Println(r.Name, r.Boat.Product.Price(0.2), r.Captain, r.FirstOfficer)
	// 	} else {
	// 		fmt.Println(r.Name, r.Boat.Product.Price(0.2))
	// 	}
	// }
	//kayak := store.NewProduct("Kayak", "Watersport", 275)
	// deal := store.NewSpecialDeal("Weekend Spesial", product, 52)
	// name, price, Price := deal.GetDetails()
	// fmt.Println(name, price, Price)
	// type OfferBundel struct {
	// 	*store.SpecialDeal
	// 	*store.Product
	// }
	// bundel := OfferBundel{
	// 	store.NewSpecialDeal("Weekend Special", kayak, 52),
	// 	store.NewProduct("Lifejacket", "Watersport", 49.54),
	// }
	// fmt.Println(bundel.Price(0)) нельзя так делать в связи с тем что функция Price
	// используется по разному в SpecialDeal and Product
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 275, 1, false),
		"Ball":  store.NewProduct("Ball", "Soccer", 19.5),
	}
	for key, p := range products {
		switch item := p.(type) {
		case store.Describe:
			fmt.Println(item.GetName(), item.GetCategory(), item.Price(0.2))
		default:
			fmt.Println(key, p.Price(0.2))
		}
	}
	//Нельзя использовать разные типы в мапе так как в Golang  нет классов
	// как в других яп , где лодка бы считалась подклассом продукта
}
