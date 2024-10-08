package main

import "fmt"

// func printPrice(product string, price, _ float64) {
// 	taxAmount := price * 0.25
// 	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
// }

// func main() {
// 	printSuppliers("Kayak", "Acme Kayak's", "Bob Boot's", "Crazy Canoes")
// 	printSuppliers("Lifejacket", "Sail safe Co")
// 	printSuppliers("Soccer ball") //для переменного типа параметра этой строки нет (если нет проверки:))
// 	names := []string{"Alan Wake", "Toby Mag", "Tomas Shelby", "Max Payne"}
// 	printSuppliers("Hat", names...) //использование переменного параметра, чтобы прописывать срез
// }

//	func printSuppliers(product string, suppliers []string) {
//		for _, supplier := range suppliers {
//			fmt.Println("product:", product, "Supplier:", supplier)
//		}
//	}
//
// Вторая реалицация без необходимости создания среза (с использованием переменного параметра (...))
//
//	func printSuppliers(product string, suppliers ...string) {
//		if suppliers == nil {
//			fmt.Println("Product:", product, "Supplier (none)")
//		} else {
//			for _, supplier := range suppliers {
//				fmt.Println("Product:", product, "Supplier:", supplier)
//			}
//		}
//	}
// func swapValues(first, second *int) {
// 	fmt.Println("Before swap", *first, *second)
// 	*first, *second = *second, *first
// 	fmt.Println("After swap:", *first, *second)
// }

// Использование указателей в функции для замены их значений не только в функции
//
//	func main() {
//		val1, val2 := 10, 20
//		fmt.Println("Before calling function:", val1, val2)
//		swapValues(&val1, &val2)
//		fmt.Println("After calling function:", val1, val2)
//	}
// func calcTax(price float64) float64 {
// 	return price + price*0.2
// }

// func Taxi(price float64) float64 {
// 	return price * 0.2
// }

//	func main() {
//		products := map[string]float64{
//			"Kayak":      275,
//			"Lifejacket": 48.95,
//		}
//		for product, price := range products {
//			fmt.Println("Product:", product, "Price with Tax:", calcTax(price), "Tax:", Taxi(price))
//		}
//	}
// func swapValues(first, second int) (int, int) {
// 	return second, first
// }

// func main() {
// 	val1, val2 := 10, 20
// 	fmt.Println("Before swap:", val1, val2)
// 	val1, val2 = swapValues(val1, val2)
// 	fmt.Println("After swap:", val1, val2)
// }

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

// func calcTotalPrice(products map[string]float64, minSpent float64) (total, tax float64) {
// 	total = minSpent
// 	for _, price := range products {
// 		if taxAmount, due := calcTax(price); due {
// 			total += taxAmount
// 			tax += taxAmount
// 		} else {
// 			total += price
// 		}
// 	}
// 	return
// }

func calcTotalPrice(products map[string]float64) (total, count float64) {
	count = 52
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	total, _ := calcTotalPrice(products)
	fmt.Println("Total:", total)
	// total1, tax1 := calcTotalPrice(products, 10)
	// fmt.Println("Total 1:", total1, "tax1:", tax1)
	// total2, tax2 := calcTotalPrice(nil, 10)
	// fmt.Println("Total 2:", total2, "tax2:", tax2)
	// for product, price := range products {
	// 	taxAmount, taxDue := calcTax(price)
	// 	if taxDue {
	// 		fmt.Println("Product:", product, "Price with tax:", taxAmount)
	// 	} else {
	// 		fmt.Println("Product:", product, "No tax due")
	// 	}
	// }
}
