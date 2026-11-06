package main

import "fmt"

// for -> only construct for looping in go lang.
func main(){

	// using as while :
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()



	// infinite While loop :
	// for {
	// 	fmt.Println(1)
	// }



	// classic for loop :
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println()


	
	// range in for loop :
	for i := range 11 {
		fmt.Println(i)
	}
}