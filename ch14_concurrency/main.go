package main

import (
	"fmt"
	//"time"
)

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println(
			"Dispatch to", details.Customer,
			":", details.Quantity,
			"x", details.Product.Name,
		)
	}
	fmt.Println("Channel has been closed")
}

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)

	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

}
