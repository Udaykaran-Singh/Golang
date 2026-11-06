package main

import "fmt"


// interface
type payment_interface interface{
	pay(amount float32) // here return val can also be given

	// interface work with only those classes, having functions >= interface
	// if any function in interface is not in class it can not be used in that

}
// open for extension but close for modification principle is vio  lates without interface



// class
type payment struct{
	gateway payment_interface
}

// class Method
func (p *payment) makePayment(amount float32){
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}




// class 1
type razorpay struct{}

// class Method 1
func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("make payment using razorpay", amount)
}




// class 2
type stripe struct{}

// class Method 2
func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe", amount)
}




func main(){
	a := razorpay{}
	new_payment := payment{
		gateway: a,
	}

	new_payment.makePayment(100)
}