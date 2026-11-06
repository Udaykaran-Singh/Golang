package main

import "fmt"

// ... means we dont know how much will we send
// ...int means only integers should be passed
func sum(nums ...int)int{
	total := 0

	// same as *agrs in python
	// gives the slice
	for _, num := range nums{
		total += num
	}

	return  total
}

func main(){
	fmt.Println(sum(1, 2, 3, 4, 5))


	// how to send all the elements of slice in variadic func
	nums := []int{1, 2, 3, 4}
	fmt.Println((sum(nums...)))
}