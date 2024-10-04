package main

import (
	"fmt"
	//"strconv"
)

func main() {
	product := "Kayak"
	for _, character := range product {
		fmt.Println("Character -", string(character))
	}
}
