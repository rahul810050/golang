package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("error agya bhai in file creation")
		return
	}
	defer file.Close()

	content := "hey there how you are doing\n"
	byt, e := io.WriteString(file, content) // it returns byte and error
	fmt.Println("byte written:", byt)
	if e != nil {
		fmt.Println("error agya bhai likhne me")
		return
	}

	// open file read content
	f, error := os.Open("text.txt")
	if error != nil {
		fmt.Println("error while opening file", error)
		return
	}

	defer f.Close()
	// create a buffer to read the file content
	buffer := make([]byte, 1024)

	for {
		n, er := f.Read(buffer)
		if er == io.EOF { // it reaches the end of the file then break
			break
		}
		if er != nil {
			fmt.Println("error while reading the file", er)
			return
		}

		// process the read content
		fmt.Println(string(buffer[:n]))
	}

	fmt.Println("successfully file created")

	// read the entire file into a byte slice 
	// ioutil.ReadFile is a convenient for scenarios where you want to read the entire content of a file into memory
	// if the size of the file is too big then there might be a problem of shortage of memory that is why we use buffer
	txt, errr := os.ReadFile("text.txt")
	if errr != nil{
		fmt.Println("error while reading the file", errr)
		return
	}
	fmt.Println(string(txt))
}
