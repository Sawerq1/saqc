package main

import (
	"fmt"
	"time"
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

func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Sent via channel 1")
		case channel2 <- p:
			fmt.Println("Sent via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}

// func main() {
// 	dispatchChannel := make(chan DispatchNotification, 100)
// 	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
// 	//var resiveOnlyChannel <-chan DispatchNotification = dispatchChannel

// 	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
// 	productChannel := make(chan *Product)
// 	go enumerateProducts(productChannel)
// 	openChannels := 2
// 	// reciveDispatches((<-chan DispatchNotification)(dispatchChannel))
// 	for {
// 		select {
// 		case details, ok := <-dispatchChannel:
// 			if ok {
// 				fmt.Println("Dispatch to", details.Customer, ":", details.Quatity, "x", details.Product.Name)
// 			} else {
// 				fmt.Println("Dispatch channel has been closed")
// 				dispatchChannel = nil
// 				openChannels--
// 			}
// 		case product, ok := <-productChannel:
// 			if ok {
// 				fmt.Println("Product:", product.Name)
// 			} else {
// 				fmt.Println("Product channel has been closed")
// 				productChannel = nil
// 				openChannels--
// 			}
// 		default:
// 			if openChannels == 0 {
// 				goto alldone
// 			}
// 			fmt.Println("-- No message ready to be recived")
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}
// alldone:
// 	fmt.Println("All values recived")
// }

func main() {
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProducts(c1, c2)
	time.Sleep(time.Second)
	for p1 := range c1 {
		fmt.Println("Channel 1 recived product:", p1.Name)
	}
	for p2 := range c2 {
		fmt.Println("Channel 2 recived product:", p2.Name)
	}
}
