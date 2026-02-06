package main

import (
	"fmt"
)

func main(){
	var name[2]string // this is how we declare array it will assign empty string to each element
	fmt.Println(name); // [  ]
	fmt.Printf("it will print the quoted string %q\n", name) // output ["" ""]

	name[0] = "cocane"
	name[1] = "rahul"
	fmt.Println(name)
	fmt.Printf("name of the persons %q\n", name) // it will print the qouted string like this ["cocane" "rahul"]

	var numbers = [5]int{1,2,3,4,5} // this is how we initialize array

	fmt.Println(numbers)
	fmt.Println("the lenght of Numbers:", len(numbers))

	var prices[5]int
	fmt.Println(prices) // it will print [0 0 0 0 0] whenever we initialize an array GOLANG assign all the element as "0"
	var trueFalse[2]bool
	fmt.Println(trueFalse) // it will print [false false] whenever we initialize an array GOLANG assign all the element as "false"
}