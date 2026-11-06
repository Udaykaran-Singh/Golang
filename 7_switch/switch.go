package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	// no need to give the brake after each case
	switch i {
	case 1:
		fmt.Println("one")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("six")
	default: // Optional
		fmt.Println("other")

	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Not Integer and string, something else :", t)
		}
	}

	whoAmI("golang")
	whoAmI(50.50)
	whoAmI(true)
	whoAmI(whoAmI)

}
