package main

import "fmt"

func main(){
	fmt.Println("learning Channel in golang")
	/*
	Channels are Go's superpower
	If goroutines are workers then channels are how they talk

	Q. what is channel??
	 -> A pipe that lets goroutines send data to each other safely
	 Goroutine A  --->  | channel |  --->  Goroutine B
	 No locks No shared mermory just send and receive

	*/
}