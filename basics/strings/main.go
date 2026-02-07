package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := "apple,orange,banana"
	parts := strings.Split(fruits, ",") // it will give a value separated by sep
	fmt.Println(parts)

	str := "one two three four two two five"
	ct := strings.Count(str, "two")
	fmt.Println(ct)

	str = "        hello world!!         "
	trimed := strings.TrimSpace(str)
	fmt.Println(trimed)

	str1, str2 := "cocane", "rahul"
	concate := strings.Join([]string{str1, str2}, " ") // it will concate strings keeping whitespace between them 
	fmt.Println(concate)
}
