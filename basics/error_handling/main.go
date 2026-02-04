package main

import (
	"fmt"
)


func divide(a,b float64) (float64, error){ // this is how we can define error output as well
	if b == 0{
		return 0, fmt.Errorf("denominator must not be zero") // error output
	}
	return a/b, nil
}

func main(){
	res, _ := divide(10, 3) // in Go the underscore (_) is used as a blank identifier. It serves as a write only variable that allows you to discard values returned by a function or to ignore specific values when you are not interested in using them
	// underscore means we are not using this error we are just ignoring it
	fmt.Println(res)
	ans, err := divide(10, 0) // we can also use other name of the variable but we have to use that variable in someway it could be in a if condition or some other way 
	if err != nil{
		fmt.Println("error handling")
	}
	fmt.Println(ans)
}