package main

import "fmt"

// func printSlice(item []int){
// 	for _, item := range item{
// 		fmt.Println(item)
// 	}
// }


// func printSlice[T any](item []T){
func printSlice[T int | string](item []T){
	for _, item := range item{
		fmt.Println(item)
	}
}

// T any ~ T interface{}        both means same thing
// but this method generaly is not good

// :- better to use comparable






// generic in struct/class
type stack[T any] struct{
	elements []T
}

func main(){
	printSlice([]int{1, 2, 3})

	names := []string{"dj", "dsifj"}
	printSlice(names)



	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)


	myStack2 := stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack2)
}