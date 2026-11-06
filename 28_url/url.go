package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to handling URLs in Go")

	// Parse the URL
	res, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", res.Scheme)
	fmt.Println("Host:", res.Host)
	fmt.Println("Path:", res.Path)
	fmt.Println("Port:", res.Port())
	fmt.Println("Raw Query:", res.RawQuery)

	qparams := res.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["watch?v"])


	for _, val := range qparams{
		fmt.Println("Param is : ", val)
	}



	// not important
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
