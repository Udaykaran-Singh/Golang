package main

import "fmt"

const city = "Bathinda"

func main(){
	const name = "Udaykaran Singh"
	fmt.Println(name, city)

	// const grouping
	const(
		port = 5000
		host = "localhost/'Udaykaran Singh'"
	)

	fmt.Println(port, host)
}