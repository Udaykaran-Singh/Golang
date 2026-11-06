package main

import "fmt"

func changeNum(num *int){
	*num = 5
	fmt.Println("in function", num)
}

func main(){
	num := 1
	changeNum(&num)

	fmt.Println("in main() :", num)
}