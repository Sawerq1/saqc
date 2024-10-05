package main

import "fmt"

func main() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Padding"
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	fmt.Println(names)
	fmt.Println(appendedNames)
}
