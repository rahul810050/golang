package main

import (
	// "basics/myutil"
	"fmt"
)

func main(){
	// fmt.Println("hey there i am learning Golang")
	// myutil.PrintHello("Hii there")

	var a int = 10
	fmt.Println(a);
	var name string = "cocane"
	fmt.Println(name)

	var b = 10
	b = 20
	fmt.Println(b)

	c := 20 // this is the short hand way of declaring a variable
	fmt.Println(c)

	var ( // this is the way to declare multi variables at once
		age = 21
		Name = "cocane"
	)
	fmt.Println(age, Name)
	const pi = 3.14
	fmt.Println(pi)

	var PublicVar = 10 // this is the public variable which can be accessed from other packages
	var privateVar = 20 // this is the private variable which can be accessed only within the package

	fmt.Println(PublicVar, privateVar)

	

	PublicFunc("Hey bro this is a Public function man and this is a new concept -> and we can use it outside of this file")
	privateFunc("hey bro this is a private function we can only use this function in this file")

}

func PublicFunc (message string){
	fmt.Println("this is the public function which can be accessed from other packages")
	fmt.Println(message)
}

func privateFunc(message string){
	fmt.Println(message)
	fmt.Println("this is the private function which can be accessed only within the package")
}