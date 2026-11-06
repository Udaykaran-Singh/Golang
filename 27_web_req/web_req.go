package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/watch?v=ru53LpdVHn4&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25"

func main(){
	fmt.Println("LCO web request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}


	fmt.Printf("Response is of type : %T\n", resp)

	defer resp.Body.Close() // caller's responsibility to close the connections


	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}