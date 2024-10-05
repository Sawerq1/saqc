package main

import "fmt"

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	someNames = append(someNames, "Boots")
	fmt.Println(someNames)
	fmt.Println("Len of someNames-", len(someNames), "Cap of someNames-", cap(someNames))
	fmt.Println(allNames)
	fmt.Println("Len of allNames-", len(allNames), "Cap of allNames-", cap(allNames))
}
