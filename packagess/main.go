package main

import (
	"fmt"
	//	cuurencyFmt "packagess/fmt"
	. "packagess/fmt"
	"packagess/store"
	"packagess/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "Watersport", 275)
	cart := cart.Cart{
		CustomerName: "Alise",
		Products:     []store.Product{*product},
	}
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
	fmt.Println("Price with $:", ToCurrency(product.Price()))
	fmt.Println(cart.CustomerName)
	fmt.Println("$", cart.GetTotal())
}
