package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	a = 27.5
	b := math.RoundToEven(a)
	fmt.Println(b)
}
