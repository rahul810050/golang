package main

import "fmt"

func main(){
	age := 21
	name := "cocane"
	height := 5.9234
	fmt.Println("age is", age, "and name is", name, "and height is", height)


	fmt.Printf("age is %d\n", age);
	fmt.Printf("name is %s\n", name);
	fmt.Printf("height is %.3f\n", height); // this will print the height with 3 decimal places
	fmt.Printf("type of age is %T\n", age); // this will print the type of age
}