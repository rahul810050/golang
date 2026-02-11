package main

import (
	"crypto/rand"
	"fmt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*"

func passGenerator(length int)(string, error){
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := range b{
		b[i] = charset[b[i] % byte(len(charset))]
	}
	return string(b), nil
}

func main(){
	pass, err := passGenerator(10)
	if err != nil {
		fmt.Println("error a gya bhai", err)
	}
	fmt.Println(pass)
}