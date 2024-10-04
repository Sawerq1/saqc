package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ { // loop with counter
		PrintHello()   // print out the messege
		PrintNumber(i) // print out the counter
	}
}

func PrintHello() {
	fmt.Println("Hello, Go")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
