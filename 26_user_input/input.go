package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	
	// enter the input 

	input, _ := reader.ReadString('\n')

	fmt.Println("your input is :", input)

	fmt.Printf("type : %T\n", input)

	num, _ := strconv.ParseFloat(input, 64)

	fmt.Println(num + 20)

}