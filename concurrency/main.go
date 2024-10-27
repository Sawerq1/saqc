package main

import (
	"fmt"
	// "time"
)

// func main() {
// 	fmt.Println("Main function started")
// 	CalcStoreTotal(Products)
// 	// time.Sleep(time.Second * 5)
// 	fmt.Println("Main function finished")
// }

// func main() {
// 	dispatchChannel := make(chan DispatchNotification, 100)
// 	go DispatchOrders(dispatchChannel)
// 	for {
// 		if details, open := <-dispatchChannel; open {
// 			fmt.Println("Dispatch to", details.Customer, ":", details.Quatity, "x", details.Product.Name)
// 		} else {
// 			fmt.Println("Channel has been closed")
// 			break
// 		}
// 	}
// }

func reciveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quatity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed.")
}

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var resiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	reciveDispatches((<-chan DispatchNotification)(dispatchChannel))
}
