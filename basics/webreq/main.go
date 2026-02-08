package main

import (
	"fmt"
	"io"
	"net/http"
)


func main(){
	fmt.Println("web request in golang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/3")
	if err != nil {
		fmt.Println("error while getting the response", err)
		return
	}
	fmt.Printf("type of the response: %T\n", res) // http.Response
	// we use defer in the purpose of if something is not in use then it should be closed
	defer res.Body.Close() // this wil close the stream on 


	// read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while reading the response", err)
		return
	}

	fmt.Printf("type of the data: %T\n", data) // byte 
	fmt.Println("data in byte form:", data) // this will print the data in byte form

	res_data := string(data) // this will convert this data into string 
	fmt.Println(res_data) // this will print the string data
}