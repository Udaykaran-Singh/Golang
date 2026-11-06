package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	
}

// order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision

	customer  // struct embedding
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

func main(){
	b :=newOrder("ufs", 9034.34, "fdg")

	fmt.Println(b)
}