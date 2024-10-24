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
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Yach", 50000, 5, true, true, "Bobby", "Charle"),
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Super Yach", 100000, 15, true, true, "Alice", "Shlapa"),
	}
	for _, r := range rentals {
		if r.IncludeCrew {
			fmt.Println(r.Name, r.Boat.Product.Price(0.2), r.Captain, r.FirstOfficer)
		} else {
			fmt.Println(r.Name, r.Boat.Product.Price(0.2))
		}
	}
}
