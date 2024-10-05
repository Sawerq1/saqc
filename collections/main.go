package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	overNames := &names
	names[0] = "Canoe"
	fmt.Println(names)
	fmt.Println(*overNames)
}
