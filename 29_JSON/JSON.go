package main

import (
	"encoding/json"
	"fmt"
)

// should start from capitals as small are treated as local
type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`       // Skip this field in JSON
	Tags     []string `json:"tags,omitempty"`
}


func main(){
	fmt.Println("Welcome to JSON video")

	// converting into valid JSON
	// EncodeJson()
	DecodeJson()

}

func EncodeJson(){
	lcoCourses := []course{
		{"reactjs bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"mern bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"angular bootcamp", 299, "learncodeonline.in", "efj123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){
	jsonDataFrowmWeb := []byte(`
	{
		"coursename" : "ReactJS Bootcamp",
		"Price" : 299,
		"website" : "LearnCodeOnline.in",
		"tags" : ["web-dev", "js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFrowmWeb)

	if checkValid{
		json.Unmarshal(jsonDataFrowmWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("MOYE MOYE")
	}


	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}   // we know key type but not the value type
	json.Unmarshal(jsonDataFrowmWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourse)


	for k, v := range myOnlineData{
		fmt.Println(k, v)
	}
}