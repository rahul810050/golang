package main

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	// "strings"

	// "io"
	"net/http"
)

type Todo struct {
    UserId int `json:userId`
    Id int `json:id`
    Title string `json:title`
    Completed bool `json:completed`
}



func performPostReq(){
    todo := Todo{
        UserId: 23,
        Title: "go to gym",
        Completed: false,
    }
    // conver the struct to json
    todoJson, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("error while converting struct into json", err)
        return
    }
    // // convert json data into string
    // todoString:= string(todoJson)
    // // reader format 
    // todoReader := strings.NewReader(todoString)

    // convert bytes into reader
    todoReader := bytes.NewReader(todoJson)

    myUrl := "https://jsonplaceholder.typicode.com/todos/"
    // it takes url, content type, body in reader format
    res, err := http.Post(myUrl, "application/json", todoReader)
    if err != nil{
        fmt.Println("error while getting response", err)
        return
    }
    defer res.Body.Close()

    // read the data
    data, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("error while reading the data", err)
        return
    }
    fmt.Println(string(data))
}

func performPutReq(){
    todo := Todo{
        UserId: 212,
        Title: "go to gym bro...go and get the gain",
        Completed: false,
    }
    // convert struct into json
    todoJson, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("error while converting struct into json", err)
        return
    }
    todoReader := bytes.NewReader(todoJson)
    myUrl := "https://jsonplaceholder.typicode.com/todos/1"

    req, err := http.NewRequest(http.MethodPut, myUrl, todoReader)
    req.Header.Set("content-type", "application/json")

    // create client
    client := http.Client{}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("error while getting response", err)
        return
    }
    defer res.Body.Close()
    data, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("error while reading the data", err)
        return
    }
    fmt.Println(string(data))
}


func performDeleteReq(){
    myUrl := "https://jsonplaceholder.typicode.com/todos/1"
    req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
    if err != nil {
        fmt.Println("error while creating delete req", err)
        return
    }
    client := http.Client{}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("error while sending the request", err)
        return
    }
    defer res.Body.Close()
    fmt.Println(res.StatusCode)
}


func main(){
    fmt.Println("learning CRUD in go")
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        fmt.Println("error agya bhai", err)
        return
    }
    defer res.Body.Close()
    if res.StatusCode != http.StatusOK {
        fmt.Println("error in while getting response", res.Status)
        return
    }
    // instead of doing this 
    // data, err := io.ReadAll(res.Body)
    // if err != nil {
    //     fmt.Println("error while reading the response", err)
    //     return
    // }
    // fmt.Println(string(data))
    // jsonData, err := json.Marshal(data)
    // if err != nil {
    //     fmt.Println("error while converting byte data into json data", err)
    //     return
    // }
    // fmt.Println(string(jsonData))

    // we can simply do this
    var todo Todo
    err = json.NewDecoder(res.Body).Decode(&todo)
    if err != nil {
        fmt.Println("error while decoding", err)
        return
    }
    fmt.Println("todo is: ",todo)
}