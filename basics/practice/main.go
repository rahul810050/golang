package main

import (
	"encoding/hex"
	"fmt"
	// "math/rand"
	"crypto/rand"
	// "time"
	// "sort"
)

type User struct {
	Name  string
	Email string
	Age   int
}


func tokenGenerator(n int)(string, error){
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil{
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func otpGenerator()(int, error){
	b := make([]byte, 2) // create byte slice with 2 allocated bytes each byte = 8 bits --> total randomness = 16 bits
	_, err := rand.Read(b) // it will put two random number from os entropy
	if err != nil {
		return 0, err
	}
	otp := int(b[0])<<8 +int(b[1]) // lets say b[0] = 213(in binary form) b[1] = 47(in binary form) ---- 213 << 8(left shift) = 213 * 256 = 54528  + 57 
	return otp % 1000000, nil // 54585 % 1000000 it will return 6 digit number
}

func passGenerator(length int)(string, error){
	// charset length is 62
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	_, err := rand.Read(b) // lets say b = [193, 44, 89, 3, 255, 101, 76, 90]
	if err != nil {
		return "", err
	}
	for i := range b{
		b[i] = charset[b[i] % byte(len(charset))] // b[0] = 193, 193 % 62 = 7 --> charset[7] = 'h' --> b[0] = 'h'
	}
	return string(b), nil // it will convert slice of bytes into string
}

func main() {
	fmt.Println("practice makes man perfect---i will become one of the best developer in the world")
	// arr := []int{}
	// var n int
	// fmt.Println("enter the nunmbers: ")
	// fmt.Scan(&n)
	// for i := 0; i < n; i++{
	// 	var el int
	// 	fmt.Scan(&el)
	// 	arr = append(arr, el)
	// }
	// sort.IntsAreSorted(arr) // it will check that is this slice sorted or not
	// sort.Ints(arr) // it will sort the slice
	// fmt.Println("the elements in the arr are: ")
	// for idx, val := range arr{
	// 	fmt.Println(idx, val)
	// }

	// lang := []string{"c++", "js", "ts", "java", "golang", "python"}
	// fmt.Println(lang) // [c++ js ts java golang python]
	// var idx int = 3
	// lang = append(lang[:idx], lang[idx+1:]...) // it will append values from idx 0 before idx then from idx+1 till the end
	// fmt.Println(lang)                          // [c++ js ts golang python]

	// lang = append(lang[:idx], lang...)
	// fmt.Println(lang)


	// cocane := User{"cocane", "cocane@gmail.com", 21}
	// fmt.Println(cocane) 
	// fmt.Printf("cocane details are: %+v\n", cocane) // this will print struct in a structured way

	// switch case 
	// old way to generate random number (Global state)
	// rand.Seed(time.Now().UnixNano())
	// diceNumber := rand.Intn(6)+1
	// fmt.Println("the value of dice:",diceNumber)

	// better way (local generator)
	// r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	// dice := r.Intn(6)+1
	// fmt.Println(dice)

	// when to use math/rand v/s crypto/rand
	// Game, dice ----> math/rand
	// password   ----> ❌
	// tokens     ----> ❌
	// security	  ----> crypto/rand


	// crypto/rand
	// this package uses or-level entropy
	// pulls randomness from 
	//  - cpu noise, mouse movement, keyboard timing, hardware RNG
	// which is unpredictable and aslo cryptographically secure

	b := make([]byte, 16) // this allocates 16 bytes
	n, err := rand.Read(b); // fills with random bytes and returns (n int, err error)
	if err != nil {
		panic(err)
	}
	fmt.Println(n," ",b);
	fmt.Println(hex.EncodeToString(b)) // it convert it to hex


	// otp generator
	otp, _ :=otpGenerator()
	fmt.Println("otp is:", otp)


	// password generator
	pass, _ := passGenerator(8)
	fmt.Println("password generator",pass)

}
