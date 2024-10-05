package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, name := range names {
		fmt.Println("Value:", name)
	}

}
