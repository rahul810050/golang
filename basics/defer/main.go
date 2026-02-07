package main

import "fmt"

func add(a,b int) int{
	return a + b
}

func main(){
	fmt.Println("before")
	data := add(5, 6)
	defer fmt.Println(data)
	defer fmt.Println("middle")
	fmt.Println("after")

	// whenever we use "defer" it acts like stack(LIFO)
	//   |  			             |
	//   |							 |
	//   |defer fmt.Println("middle")| first this will be the output after execution all the like except defer keyword line
	//   |defer fmt.Println(data)	 | second this will be the out put
	//   |___________________________|
}