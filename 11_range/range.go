package main

import "fmt"

func main(){
	nums := []int{7, 8, 9}

	for i := 0; i < len(nums); i++{
		fmt.Println(nums[i])
	}

	fmt.Println()

	// a gives the index
	// num gives the value
	for a, num := range nums{
		fmt.Println(a, num)
	}
	fmt.Println()

	// use in case of map
	m := map[int]int{1:1, 2:2}

	for k, v := range m{
		fmt.Println(k, v)
	}
	fmt.Println()

	// for string -> unicode
	for i, c := range "hello how are you?"{
		fmt.Println(i, c)
	}

}