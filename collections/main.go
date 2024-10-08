package main

import (
	"fmt"
	//"strconv"
	//"reflect"
	//"sort"
)

func main() {
	var price string = "€48.95"
	// var currency string = string(price[0])
	// var ammountString string = string(price[1:])
	// amount, parseErr := strconv.ParseFloat(ammountString, 64)
	// fmt.Println("Currency:", currency)
	// fmt.Println("Len:", len(price))
	// if parseErr == nil {
	// 	fmt.Println("Ammount:", amount)
	// } else {
	// 	fmt.Println("Parse error:", parseErr)
	// }
	for index, char := range price {
		fmt.Println("Index:", index, "Char:", char, "StringChar:", string(char))
	}
}
