package main

import (
	"fmt"
	"time"
)


func sayHello(){
	fmt.Println("hello there")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("function ended")
}

func sayHi(){
	fmt.Println("hi there")
	time.Sleep(2000 * time.Millisecond)
}

func main(){
	fmt.Println("learning Go-Routines")
	go sayHello()
	go sayHi()

	time.Sleep(3000 * time.Millisecond)
}