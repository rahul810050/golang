package main

import "fmt"

type Number interface {
	int | float64 | uint
}

type GenericSlice[T any] []T // defining my own generic slice

type GenericStruct[T any] struct{
	values T
}

func (g GenericSlice[T]) Print(){
	for _, val := range g{
		fmt.Println(val)
	}
}

func add[T Number](x T, y T) T{ // generics is something like an argument of any function can be int as well as float and if we make the add function with only int and if there are arguments which are float type then it will throw error soo for that we use GENERICS
	return x + y
}

func getValues[K comparable, V any](mp map[K]V) []V{
	val := []V{}
	for _, value := range mp{
		val = append(val, value)
	}
	return val
}

func main(){
	fmt.Println(add(1,2))
	mp := map[int]string{1: "hello", 2: "there"}
	val := getValues(mp)
	fmt.Println(val)

	// generic type
	g := GenericSlice[int]{1,2,3,4}
	fmt.Println(g)
	g.Print()
}