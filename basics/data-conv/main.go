package main

import (
	"fmt"
	"strconv"
)


func main(){
	num := 123
	var float float64 = float64(num)
	fmt.Printf("value of float %.3f\n", float);
	fmt.Printf("type of float %T\n", float);

	num_str := "1234"
	num_int, _ := strconv.Atoi(num_str)
	fmt.Printf("type of num_int %T\n", num_int)


	str := "3.14"
	num_float, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("type of num_float %T\n", num_float);
}