package main

import (
	"fmt"
	//"strconv"
)

type calcFunc func(float64) float64

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.96,
	}
	for product, value := range products {
		printPrice(product, value, selectCalculator(value))
	}
}
