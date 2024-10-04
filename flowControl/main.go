package main

import (
	"fmt"
	// "strconv"
)

func main() {
	count := 0
target:
	fmt.Println("Counter -", count)
	count++
	if count < 5 {
		goto target
	}
}
