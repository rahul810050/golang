package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var name string
	fmt.Println("hey what is your name??:")
	fmt.Scan(&name) // that is how we take input--- keep in mind that fmt.Scan reads until the first whitespace characters
	fmt.Printf("hey there!! very nice to meet you, %s",name)

	// bufio reader 
	reader := bufio.NewReader(os.Stdin) // it creates a new buffer reader that reads from the standard input (os.Stdin)
	name, _ = reader.ReadString('\n') // reads a line from the input until it encounters a new line
	fmt.Println("Hii ", name)
}