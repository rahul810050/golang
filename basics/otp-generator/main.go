package main

import (
	"crypto/rand"
	"fmt"
)


func otpGenerator()(int, error){
	b := make([]byte, 2)
	_, err := rand.Read(b);
	if err != nil {
		return 0, err
	}
	otp := int(b[0])<<8 + int(b[1])
	return otp % 1000000, nil
}

func main(){
	otp, err := otpGenerator()
	if err != nil {
		fmt.Println("error a gya bhai", err)
	}
	fmt.Println(otp)
}