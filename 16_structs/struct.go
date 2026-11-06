package main

import (
	"fmt"
	"time"
)

// order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}


func newOrder(id string, amount float32, status string) *order{
	// initaial setup
	s := order{
		id: id,
		amount: amount,
		status: status,
	}

	return &s
}

// creating class functions/Methods
// aka receiver type
func (self *order) changeStatus(status string){
	self.status = status
}


func main(){
	a := order{
		id : "1",
		amount: 50.00,
		status: "recieved",
	}

	a.createdAt = time.Now()

	fmt.Println(a)
	fmt.Println()

	a.changeStatus("confirmed")

	fmt.Println(a)
	fmt.Println()


	b := newOrder("1", 30.5, "recieved")
	fmt.Println(*b) 
}