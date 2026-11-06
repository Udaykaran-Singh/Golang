package main

import (
	"fmt"
	"slices"
)

// slice -> Dynamic array just like vectors
// most used construct in go
func main(){

	// uninitialized slice is nil
	var nums[] int
	fmt.Println("is nil ? :-", nums == nil)
	fmt.Println("len of nums now :", len(nums))
	fmt.Println()



	// type, size, capacity
	// make()
	var arr = make([]int, 1, 1)
	fmt.Println(arr)
	fmt.Println(arr == nil)
	fmt.Println("capacity :",cap(arr))

	// add elements
	arr = append(arr, 1)
	// arr = append(arr, 2)

	fmt.Println("capacity :",cap(arr))
	fmt.Println(len(arr))
	fmt.Println()



	// make slice way 3 :-
	arr1 := []int{}
	fmt.Println(arr1)
	fmt.Println("capacity :",cap(arr1))
	fmt.Println(len(arr1))
	fmt.Println()



	// copy syntax
	source := []int{1, 2, 3}
	// source = append(source, 1)

	// making destination same size as source
	destination := make([]int, len(source))

	copy(destination, source)
	fmt.Println(destination, source)




	fmt.Println(source[:2])
	// source = append(source, 3)
	fmt.Println(slices.Equal(source, destination))




	// 2d
	var nums10 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums10)
	
}