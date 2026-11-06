package main

import "fmt"

func counter() func()int {
	// always presist as below given local is passed as alice
	var count int = 0;

	// return function and will make that var as an alice
	return func()int { 
		count += 1
		return count
	}
}

// closure:-

// give a variable function (define declare but not still used)
// then this variable will behave as an alias name to the function name

func main(){
	inc := counter()
	fmt.Println(inc())
	fmt.Println(inc())

}