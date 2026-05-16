package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
		fmt.Fprint(res, "method is not supported", http.StatusNotFound)
		return
	}
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(res, "POST request successful", http.StatusCreated)
	name := req.FormValue("name")
	age := req.FormValue("age")
	fmt.Print(name, " ",age)
	fmt.Fprintf(res, "Name: %s", name)
	fmt.Fprintf(res, "Age: %s", age)
}

func helloHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/hello"{
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(res, "hello there!!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // this will look for the index.html static page
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// err := http.ListenAndServe(":3000", nil)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("error while listening on port 3000")
		log.Fatal(err)
	}
}