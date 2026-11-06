package main

import "fmt"

func main(){
	// creating map

	// [key type] value type
	m := make(map[string]int)

	fmt.Println()
	m["Udaykaran"] = 100
	m["Manharan"] = 100
	m["Gurjot"] = 100
	fmt.Println(m)
	fmt.Println(m["Udaykaran"])
	// if key does not exist it return zero value


	// empty the map
	clear(m)
	fmt.Println(m)
	fmt.Println()



	// shorthand map
	ma := map[int]int{/* can also declear here initial value*/}
	ma[1] = 2
	ma[3] = 4
	fmt.Println(ma)
	fmt.Println()

	// a tells value at key
	// ok tell key present
	// if dont want to use a, can use '_' 
	// _ can not be used to print then 
	a, ok := ma[5]
	fmt.Println(a, ok)
}