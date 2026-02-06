package main

import "fmt"

func main(){
	grades := make(map[string]int) // this is how we define map in golang

	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		name, marks := "", 0
		fmt.Scan(&name, &marks)
		grades[name] = marks
	}

	delete(grades, "rahul") // this is how we delete the key from the map

	// checking a key exists or not 
	nam, mark := grades["rahul"]
	if mark == false {
		fmt.Println("name does not exists")
	} else {
		fmt.Println("name exists", nam)
	}


	person := map[string]int {
		"rahul" : 50,
		"bob": 70,
	}

	fmt.Println(person["rahul"])
}