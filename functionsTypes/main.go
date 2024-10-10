package main

import (
	"fmt"
	//"strconv"
)

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.96,
	}
	for product, value := range products {
		var calcFunc func(float64) float64
		if value > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		finalPrice := calcFunc(value)
		fmt.Println("Product:", product, "Price:", finalPrice)
	}
}
