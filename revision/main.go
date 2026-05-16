package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
)

type User struct {
	Name    string
	Age     int
	IsAdult bool
}

func add(a int, b int) int {
	return a + b
}

func main() {
	// revision of basics
	name := "cocane"
	fmt.Printf("hey there very nice to meet you, %s", name)

	sum := add(5, 10)
	fmt.Printf("\nThe sum of 5 and 10 is: %d", sum)

	user1 := User{
		Name:    "cocane",
		Age:     21,
		IsAdult: true,
	}
	fmt.Printf("\nUser data: %+v\n", user1)

	jsonData, err := json.Marshal(user1) // this converts the user into JSON
	if err != nil {
		fmt.Println("error a gya bhai", err)
	}
	fmt.Println("this will print the byte jsondata ", jsonData)
	userString := string(jsonData) // this converts JSON into String
	fmt.Println("user in string format: ", userString)

	x := "1234"
	y, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("error a gya bhai", err)
	}
	fmt.Println("the conversion will look like this: ", y)
	fmt.Printf("%T", y)

	v, err := strconv.ParseInt(x, 10, 0) // ParseInt function takes three argument(string, base like decimal/binary, bitSize)
	fmt.Println(v)

	age := 19
	switch {
	case age < 10:
		fmt.Println("child")
	case age > 12 && age < 18:
		fmt.Println("teenager")
	case age >= 18:
		fmt.Println("adult")
	}

	switch status := getStatusCode(); status {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("NOT FOUND")
		fallthrough // this just force the next case to print
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("Unknown status: ", status)
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	str := "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	// using range
	for _, ch := range str {
		fmt.Printf("%c", ch)
	}

	// array

	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr)

	arr1 := [...][2]int{{1, 2}, {3, 4}}
	arr1[0] = [2]int{2, 1} // we have to recreate an array while replacing with an array
	fmt.Println(arr1)
	// iterate over nested array
	for _, nested := range arr1 {
		for _, val := range nested {
			fmt.Println(val)
		}
	}
	testArr(arr);

	// slices
	arr2 := [5]int{1,2,3,4,5}
	sl := arr2[1:] // it will start from the index 1 till the last index
	fmt.Println(sl)

	sl1 := []string{"hello", "world"}
	fmt.Println(sl1)
	fmt.Printf("%T", sl1)

	// sl2 := []int{}
	// for i := 0; i < 5; i++{
	// 	var val int
	// 	fmt.Scan(&val)
	// 	sl2 = append(sl2, val)
	// 	fmt.Println(sl2, len(sl2), cap(sl2))
	// }
	// sll := make([]string, 10) // that is how we create slice using make function
	sl3 := []string{"hi", "there"}
	testSlice(sl3) // when we pass slice as an argument to any function it gets passed as a reference not as copy
	fmt.Println(sl3)

	// Map
	// var mp map[string]int = map[string]int{} // if we use var
	// mp := make(map[int]string) // using make function we dont have to put a curly braces at the last
	// mp1 := map[int]string{}
	// for i := 0; i < 3; i++{
	// 	var x string
	// 	fmt.Scan(&x)
	// 	mp1[i] = x
	// }
	// fmt.Println(mp1)

	// complex map
	mp2 := map[int][]int{1:{1,2,3,4}}
	mp2[2] = []int{5,6,7}
	delete(mp2, 2) // delete function takes two argument 1-> map and 2-> the key that value we want to delete
	fmt.Println(mp2)

	mp3 := map[string]int{}
	mp3["a"] = 2
	val, ok := mp3["a"] // to check the key-value exists or not, it returns two things 1-> value and 2-> ok which is a boolean the key actually present or not
	fmt.Println(val, ok)


	// functions
	hello := callFunc(doubleVal)
	fmt.Println(hello)

	f1 := getFunc("hello")
	value := f1(" world")
	fmt.Println(value)
	
	// variadic func
	sm := summation(1,2,3)
	fmt.Println(sm)

	sm1 := summation1([]int{1,2,3,4,5}...) // when we pass a slice then we need to break it "..." breaks each element into individual integer
	fmt.Println(sm1)

	// struct
	var p1 Person = Person{name: "cocane", age: 22}
	naam := p1.getName() // we send a copy of Person reference
	fmt.Println(naam)
	p1.setName("rahul") // we send a copy of Person not reference
	fmt.Println(p1)

	p2 := Person{
		name: "cocane", 
		age: 22, 
		favSports: []Sports{
			{
				name: "football", 
				position: "Full Back",
			},
			{
				name: "Moto GP",
				position: "rider",
			},
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.favSports[0].name, p2.favSports[0].position)
	fmt.Println(p2.favSports[1].name, p2.favSports[1].position)

	// interface
	var s Shape = Triangle{1,2,3}
	allSides := s.getSides()
	fmt.Println(s.getPerimeter())
	fmt.Println(s.getSides())
	fmt.Println(allSides)

	var s1 []Shape = []Shape{Triangle{1,2,3}, Square{10}} // we can create slice of different shapes
	sideT := s1[0].getPerimeter()
	sideS := s1[1].getPerimeter()
	fmt.Println(sideT, sideS)

	perimeter := uint(0)
	for _, k := range s1{
		perimeter += k.getPerimeter()
	}
	fmt.Println(perimeter)

	// error handling
	defer defferedFunc() // defer keyword is similar like finally keyword in JavaScript it will run at the last of the function
	// panic("this caused the crash") // whatever we write after panic function will not be executed
	// fmt.Println("hello")

	div, err := division(3,0)
	if err != nil {
		fmt.Println("error agya bhai",err)
	}
	fmt.Println("division of 2 numbers",div)

	// GENERICS
	fmt.Println(addition(1,2))

	// pointer
	a := 10
	pointer(&a)
	fmt.Println(a)

	b := 20
	v1 := &[]*int{&a, &b}
	test(v1)
}

// pointers
func test(pointerSlice *[]*int){
	values := *pointerSlice
	for _, value := range values{
		fmt.Println(*value)
	}
}
func pointer(x *int){
	*x = 100
}

// GENERICS
func addition[T int | float64 | uint](x T, y T) T{ // generics is something like an argument of any function can be either int type or float type and if we make the add function with only int and if there are arguments which are float type then it will throw error soo for that we use GENERICS
	return x + y
}

func defferedFunc(){
	fmt.Println("defer")
	r := recover() // recover actually catches the error/crash when a panic happens and store it inside the variable "r" ---it just tells that hey dont crash the program, i am handling the error myself you dont need to crash you can let me continue to execute 
	fmt.Println(r)
}

func divide(a int, b int) int{
	return a / b
}

func division(a int, b int) (int, error){
	if b == 0 {
		return 0, errors.New("division by 0") // we can handle error using this errors package
	}
	return a / b, nil
}

type Shape interface{
	getPerimeter() uint
	getSides() []uint
}

type Triangle struct{
	a uint
	b uint
	c uint
}

type Square struct {
	width uint
}

func (s Square) getPerimeter() uint{
	return 4 * s.width
}

func (s Square) getSides() []uint{
	return []uint{s.width, s.width, s.width, s.width}
}

func (t Triangle) getPerimeter() uint{
	return t.a + t.b + t.c
}

func (t Triangle) getSides() []uint {
	return []uint{t.a, t.b, t.c}
}

type Sports struct {
	name string
	position string
}

type Person struct {
	name string
	age uint
	favSports []Sports
}

func (p Person) getName() string {
	return p.name
}

func (p Person) setName(newName string){
	p.name = newName
	fmt.Println(p)
}


// variadic function
func summation(num ...int) int{ // it takes the argument as slice
	s := 0
	for _, val := range num{
		s += val
	}
	return s
}

func summation1(num ...int) (s int){ // named return values we dont have redeclare the variable s inside the func and dont have to return the value explicitly we just have to write return and the function will automatically return s variable
	for _, val := range num{
		s += val
	}
	return
}

func getFunc(str string) func(string) string {
	return func(str2 string) string{
		return str + str2
	}
}

func doubleVal(n int) int {
	return 2 * n
}

func callFunc(callable func(int) int) int{
	return callable(10)
}

func testSlice(arr []string){
	arr[0] = "changed this"
}

func testArr(arr [2][2]int) {
	arr[0] = [2]int{4,5}
}

func getStatusCode() int {
	return 404
}
