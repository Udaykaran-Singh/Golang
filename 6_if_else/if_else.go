package main

import "fmt"

func main(){
	age := 21

	if age >= 18 {
		fmt.Println("Person is an Adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a Kid")
	}

	//  logical operators ||, && also work here



	// we can construct var in if statement
	if age := 20; age >= 18 {
		fmt.Println("Adult")
	}


}