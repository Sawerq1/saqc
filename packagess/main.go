package main

import (
	"fmt"
	//	cuurencyFmt "packagess/fmt"
	_ "packagess/data"
	cus "packagess/fmt"
	"packagess/store"

	"packagess/store/cart"
	//"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersport", 275)
	cart := cart.Cart{
		CustomerName: "Alise",
		Products:     []store.Product{*product},
	}
	//color.Green("Name:" + cart.CustomerName)
	//color.Cyan("Total:" + cus.ToCurrency(cart.GetTotal()))
	// fmt.Println("Name:", product.Name)
	// fmt.Println("Category:", product.Category)
	// fmt.Println("Price:", product.Price())
	// fmt.Println("Price with $:", cus.ToCurrency(product.Price()))
	fmt.Println(cart.CustomerName)
	fmt.Println(cus.ToCurrency(cart.GetTotal()))
}
