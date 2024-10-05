package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := names == moreNames
	fmt.Println("Comprasion:", same)
}
