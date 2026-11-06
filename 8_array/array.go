package main

import "fmt"

func main(){
	var nums[4] int

	nums[0] = 38

	// array length
	fmt.Println("length :", len(nums))

	fmt.Println(nums[0])
	
	fmt.Println(nums)


	// short hand
	arr := [3] int{1, 2, 3}
	fmt.Println(arr)


	// 2d
	arr1 := [2][2]int{{2, 3}, {4, 5}}
	fmt.Println(arr1)
}
