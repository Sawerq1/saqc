package main

import (
	"fmt"
)

type CategoryCountMessege struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan<- CategoryCountMessege) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessege{TerminalError: arg}
			close(outChan)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessege{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {
	category := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessege)
	go processCategories(category, channel)
	for messege := range channel {
		if messege.TerminalError == nil {
			fmt.Println("Category:", messege.Category, "Total:", messege.Count)
		} else {
			fmt.Println("Terminal Error occured")
		}
	}
	// for message := range channel {
	// 	if message.CategoryError == nil {
	// 		fmt.Println(message.Category, "Total price", ToCurrency(message.Total))
	// 	} else {
	// 		panic(message.CategoryError)
	// 		//fmt.Println(message.Category, "no such catalog")
	// 	}
	// }
	// for _, cat := range category {
	// 	total, err := Products.TotalPrice(cat)
	// 	if err == nil {
	// 		fmt.Println(cat, "Total price", ToCurrency(total))
	// 	} else {
	// 		fmt.Println(cat, "no such catalog")
	// 	}
	// }
}
