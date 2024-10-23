package data

import "fmt"

func init() {
	fmt.Println("function init invoked")
}

func GetData() []string {
	return []string{"Kayak", "Lifejacket", "Soccer ball", "Paddle"}
}
