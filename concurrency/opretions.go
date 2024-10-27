package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to recover from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println(len(channel), cap(channel))
		fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))
		value := <-channel
		fmt.Println("-- channell read complete", value)
		storeTotal += value
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		//fmt.Println(p.Category, "product:", p.Name)
		total += p.Price
		// time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	// fmt.Println(category, "Subtotal:", ToCurrency(total))
	fmt.Println(category, "channel send complete")
	resultChannel <- total
}
