package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func hashPassword(password string)(string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hash, password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main(){
	var password string
	fmt.Scan(&password)
	// hash the password
	hash_pass, err := hashPassword(password)
	if err != nil {
		fmt.Println("error a gya bhai", err)
		return
	}
	// check the if the pass is correct or not
	yes := checkPassword(hash_pass, password)
	if !yes {
		fmt.Println("password incorrect!! please provide the correct password")
		return
	}
	fmt.Println("password is correct!! redirecting to dashboard")
}