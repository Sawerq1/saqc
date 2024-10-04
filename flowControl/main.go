package main

import (
	"fmt"
	// "strconv"
)

func main() {
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is more or equally 3 but less 7")
		default:
			fmt.Println(counter, "is more or equally 7")
		}
	}
}
