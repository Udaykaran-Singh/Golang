package main

import "fmt"

// enumerated types
type OrderStatus int

const(
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrdersStatus(status OrderStatus){
	fmt.Println("changing order status to", status)
}

func main(){
	changeOrdersStatus(Prepared)
}