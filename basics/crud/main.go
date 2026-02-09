package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	// "io"
	"net/http"
)

type Todo struct{
    UserId int `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

func getRequest(){
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
    if err != nil {
        fmt.Println("error while getting the response", err)
        return
    }
    defer res.Body.Close()

    // either we can do this --> read all the response using io.ReadAll function and then convert this to json using Marshall 
    // byte_res, err := io.ReadAll(res.Body)
    // if err != nil {
    //     fmt.Println("error while reading the response", err)
    //     return 
    // }
    // res_byte, err := json.Marshal(byte_res)
    // if err != nil {
    //     fmt.Println("error while converting the response into json", err)
    //     return
    // }
    // res_data := string(res_byte)
    // fmt.Println("response data:", res_data)


    // or we can use one single function json.NewDecoder to decode the response
    var todo Todo
    err = json.NewDecoder(res.Body).Decode(&todo)
    if err != nil {
        fmt.Println("error while decoding the response using new decoder", err)
        return
    }
    fmt.Println(todo)
}

func postRequest(){
    todo := Todo{
        UserId: 99,
        Id: 99,
        Title: "learn golang",
        Completed: false,
    }
    // convert the todo struct into JSON
    json_data, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("error while mashalling", err)
        return
    }
    // conver the json data into string
    json_string := string(json_data)
    // convert the json data into reader
    json_reader := strings.NewReader(json_string)

    myUrl := "https://jsonplaceholder.typicode.com/todos"
    // post method
    res, err := http.Post(myUrl, "application/json", json_reader)
    if err != nil {
        fmt.Println("error while getting the response", err)
        return
    }
    defer res.Body.Close()
    // check the statuscode
    if res.StatusCode != http.StatusOK{
        fmt.Println("error while getting the res giving status code",res.Status)
        return
    }
    // read the data
    res_data, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("error while reading the response", err)
        return
    }
    // convert the data into string
    res_data_string := string(res_data)
    fmt.Println(res_data_string)
}

func putRequest(){
    todo := Todo{
        UserId: 892897294378,
        Title: "hey there",
        Completed: true,
    }

    // convert the todo into json using marshall
    json_todo, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("error while converting the todo into json", err)
        return
    }
    /*
    // convert the json into string
    json_string := string(json_todo)
    // convert the json_string into reader 
    json_reader := strings.NewReader(json_string)
    */

    json_reader := bytes.NewReader(json_todo)

    // url
    MyUrl := "https://jsonplaceholder.typicode.com/todos/1"
    // create request 
    req, err := http.NewRequest(http.MethodPut, MyUrl, json_reader)
    if err != nil {
        fmt.Println("error while creating the request", err)
        return
    }

    // set content-type(Header)
    req.Header.Set("Content-Type", "application/json")

    // create the client
    client := http.Client{
        Timeout: 10 * time.Second, // timeout 
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("error while sending the requst", err)
        return
    }
    defer res.Body.Close()

    if res.StatusCode < 200 || res.StatusCode >= 300{
        fmt.Println("request failed", res.Status)
        return
    }
    // read the response
    data, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("error while reading the response", err)
        return
    }
    res_data := string(data)
    fmt.Println("response data:", res_data)
}


func deleteRequest(){
    myUrl := "https://jsonplaceholder.typicode.com/todos/2"

    // create request 
    req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
    if err != nil {
        fmt.Println("error while creating the request", err)
        return
    }
    // create the client
    client := http.Client{
        Timeout: 10 * time.Second,
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("error while getting the response", err)
        return
    }
    if res.StatusCode < 200 || res.StatusCode >= 300 {
        fmt.Println("request failed", err)
        return
    }
    // close the request
    defer res.Body.Close()
    
    // read the response
    res_data, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("error while reading the data", err)
        return
    }
    fmt.Println("status code:", res.Status)
    fmt.Println(string(res_data))
}

func main(){
	fmt.Println("learning CRUD in golang")
    // getRequest()
    // postRequest()
    // putRequest()
    deleteRequest()
}