package main

import "fmt"

// can also write as :-
// add(a, b int)int
func add(a int, b int) int{
	return a+b
}


// in go we can return multiple values
func get()(int, string, bool){
	return 1, "uday", true
}


// sending function as argument in another function
func p(fn func(a int)int) int{
	return fn(100)
}

func main(){
	fmt.Println(add(2, 3))
	fmt.Println()
	// a, _, _ := get()
	a, b, c := get()
	// fmt.Println(get())
	fmt.Println(a, b, c)


	fn := func(a int)int{
		return 2
	}

	fmt.Println(p(fn))
}