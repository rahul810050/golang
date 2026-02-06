package main

import (
	"fmt"
)

func main(){
	fmt.Println("if else statement in golang")
	x := 10

	if x > 20{
		fmt.Println("x is greate than 20")
	} else if x > 15 {
		fmt.Println("x is greater than 15")
	} else {
		fmt.Println("x is lower than 15")
	}

	y := 10

	if x > 5 && (y > 5 || y < 20){
		fmt.Println("hey there how are you??")
	} else {
		fmt.Println("i know you should be fine")
	}
}