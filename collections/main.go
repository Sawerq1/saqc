package main

import (
	"fmt"
	//"reflect"
	"sort"
)

func main() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	fmt.Println("map size:", len(products))
	//delete(products, "Hat")
	if _, ok := products["Hat"]; ok {
		fmt.Println("stored")
	} else {
		fmt.Println("not stored")
	}
	fmt.Printf("\n\n")
	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Printf("\n\n\nSorted map->\n")
	keys := make([]string, 0, len(products))
	for key := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}
}
