// package main

// import (
// 	"fmt"
// )

// // func getProductName(index int) (name string, err error) {
// // 	if len(Products) > index {
// // 		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
// // 	} else {
// // 		err = fmt.Errorf("Error for index: %v", index)
// // 	}
// // 	return
// // }

// // func main() {
// // 	// fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
// // 	// fmt.Print("Product: ", Kayak.Name, " Price: ", Kayak.Price, "\n")
// // 	// fmt.Printf("Product: %v Price: $%4.2f\n", Kayak.Name, Kayak.Price)
// // 	name, _ := getProductName(1)
// // 	_, err := getProductName(52)
// // 	fmt.Printf("%v\n%v\n", name, err.Error())
// // }

// func Printfln(template string, values ...interface{}) {
// 	fmt.Printf(template+"\n", values...)
// }

// func main() {
// 	Printfln("Value: %v", Kayak)
// 	Printfln("Value with fields: %+v", Kayak)
// 	Printfln("Go sintax:%#v", Kayak)
// 	Printfln("Type:%T", Kayak)
// }
