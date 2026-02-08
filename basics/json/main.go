package main

import (
	"encoding/json"
	"fmt"
)


type User struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	IsAdult bool  `json:"is_adult"`
}

func main(){
	fmt.Println("json(javascript object notation) in Golang")
	user1 := User{
		Name: "cocane",
		Age: 21,
		IsAdult: true,
	}
	fmt.Println("user data:", user1)

	// convert user1 into JSON encoding which is also called (Marshalling)
	jsonData, err := json.Marshal(user1) // this will convert the user1 struct into JSOn
	if err != nil {
		fmt.Println("error while marshalling", err)
		return
	}
	fmt.Printf("type of jsonData: %T\n", jsonData) // byte/[]uint8
	fmt.Println(jsonData) // it will print the json data in byte format
	data := string(jsonData) // convert the json data into string 
	fmt.Println("user1 in json format", data) 


	// JSON Decoding (Unmarshalling)
	var user1Data User
	e := json.Unmarshal(jsonData, &user1Data) 
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. 
	// If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	if e != nil {
		fmt.Println("error while unmarshalling", e)
		return
	}
	fmt.Printf("type of the user1Data: %T\n", user1Data)
	fmt.Println(user1Data)
}