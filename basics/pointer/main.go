package main

import (
	"fmt"
)

func modifyValuebyRef(a *int) int{
	*a = *a + 20 // *a -> 10  it will change value of "val" 
	return *a
}

func main() {
	num := 2
	ptr := &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr) // * before the pointer will print the value in that address

	var pointer *int
	if pointer == nil {
		fmt.Println("pointer is not assign")
	}

	val := 10
	sum := modifyValuebyRef(&val)
	fmt.Println(sum)
	fmt.Println(val)
}
