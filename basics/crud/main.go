package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
)

type Todo struct{
    UserId int `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

func main(){
	fmt.Println("learning CRUD in golang")
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1");
    if err != nil {
        fmt.Println("error while getting the todo from server", err)
        return
    }
    // close the stream 
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK{
        fmt.Println("error while getting the response", res.Status)
        return
    }

    // byte_data, err := io.ReadAll(res.Body)
    // if err != nil {
    //     fmt.Println("error while reading the response", err);
    //     return
    // }
    // res_data := string(byte_data)
    // fmt.Println(res_data)

    var todo Todo
    // it decodes the response into this Todo format
    err = json.NewDecoder(res.Body).Decode(&todo)
    if err != nil {
        fmt.Println("error while decoding the response", err)
        return
    }
    fmt.Println(todo)
}