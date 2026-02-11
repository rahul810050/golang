package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)


func tokenGenerator(length int)(string, error){
	b := make([]byte, length)
	_, err := rand.Read(b);
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func main(){
	fmt.Println("token generator in golang")
	token, err := tokenGenerator(16)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}