package main

import (
	"fmt"
	"time"
)

func run1(){
	time.Sleep(2 * time.Second)
	fmt.Println("this is run 1")
}

func run2(){
	time.Sleep(4 * time.Second)
	fmt.Println("this is run 2")
}

func run3(){
	time.Sleep(6 * time.Second)
	fmt.Println("this is run 3")
}

func add(x int, y int, ch chan<- int, delay int)  {
	time.Sleep(time.Duration(delay) * time.Second)
	// fmt.Println(x + y)
	ch <- x + y // this is how we can send the value through channel
	// if we comment the above line then it will run and then will give a deadlock error because it will not return anything through channel
}

func main(){ // this is the main go routine function
	// we are creating 3 go routine functions which are running in the background
	// go run1()
	// go run2()
	// go run3()
	// time.Sleep(7 * time.Second) // this will stop the cursor here for few seconds while the other goroutines are being executed
	// the main go routine function comes here and print done and there is no more code being executed in this main function soo our program is going to stop because the main goroutine has exited
	fmt.Println("done")

	// x := go add(1,2) // whenever we use go syntax while calling a function if it returns anything we cant store it like this because it will not wait for it to execute rather it will jump into the next line and execute it
	// this is where a concept comes into the picture is called channel which is a special way that we can pass values between different goroutines
	// and we are allowed to send values on the channel and wait for the value to be received
	// soo we cant actually get a return value from a goroutine instead we create and pass something known as channel 
	// channel allows us to synchronize and wait for different values to be returned
	/*
	// general channel
	// soo we cant actually get a return value from a goroutine instead we create and pass something known as channel 
	// channel allows us to synchronize and wait for different values to be returned
	// ch := make(chan int) // this is how we make channel 
	// go add(1,2,ch)
	// sum := <-ch // this is how we receive the value 
	// fmt.Println(sum)

	// ch := make(chan int)
	// go add(3, 10, ch)
	// x := <- ch
	// go add(3, 5, ch)
	// x = <- ch
	// go add(2, 5, ch)
	// x = <- ch
	// go add(10, 5, ch)
	// x = <- ch
	// fmt.Println(x)

	ch := make(chan int)
	ch2 := make(chan int)
	go add(2, 4, ch, 4)
	go add(1, 2, ch2, 2)
	// x := <- ch
	// y := <- ch2
	select {
	case x := <- ch:
		fmt.Println(x)
	case y := <- ch2:
		fmt.Println(y)
	}
	// fmt.Println(x, y)
	*/

	// directional channel
	// in the argument if (ch <-chan int) -> this means this function can only receive only channel and (ch chan<- int) - means this function can only send only channel
}