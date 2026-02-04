package main

import (
	"fmt"
)

func simpleFunc() {
	fmt.Println("hey there it is a simple function")
}

func add(a, b int) (res int) { // if
	res = a + b
	return // as we declare that res will be the output that is why even if we dont write the "res" after return still it will return res because we have defined
}

func concate(s1, s2 string) string{
	return s1 + s2
}

func main() {
	simpleFunc()
	ans := add(2, 2)
	fmt.Println(ans)

	// concatinate
	str := concate("Rahul", "Naskar")
	fmt.Println(str)
}
