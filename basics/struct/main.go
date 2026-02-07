package main

import "fmt"

type User struct {
	firstname string
	lastname  string
	username  string
	email     string
	password  string
	age       int
}



type Person struct {
	Person_Name PersonName
	Person_Contact PersonContact
	Person_Address PersonAddress 
}


type PersonName struct{
	firstname string
	lastname string
}

type PersonContact struct{
	Email string
	Mobile string
}

type PersonAddress struct{
	Home int
	Area string
	State string
	City string
}

func main() {
	var cocane User
	fmt.Println(cocane)
	// one way to assign values to struct variables 
	cocane.firstname = "rahul"
	cocane.lastname = "naskar"
	cocane.username = "cocane"
	cocane.email = "cocane@gmail.com"
	cocane.password = "random123"
	cocane.age = 21
	fmt.Println(cocane)
	fmt.Println(cocane.username) // cocane

	// 2nd way to assign values to struct variable
	sparsh := User{
		firstname: "sparsh",
		lastname: "sinha",
		username: "sparsh123",
		email: "sparsh@gmail.com",
		password: "random123",
		age: 19,
	}
	fmt.Println(sparsh)

	// 3rd way to declare and assign value to struct variables
	var manish = new(User)
	manish.firstname = "manish"
	manish.lastname = "kumar"
	manish.username = "manishdalla"
	manish.email = "manish@gmail.com"
	manish.password = "random123"
	manish.age = 24
	fmt.Println(manish)
}
