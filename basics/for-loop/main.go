package main

import "fmt"


func main(){
	// names := []string{}

	// var n int
	// fmt.Scan(&n);
	// for i := 0; i < n; i++{
	// 	var str string
	// 	fmt.Scan(&str)
	// 	names = append(names, str)
	// }

	// for _, name := range names{
	// 	fmt.Println(name)
	// }

	txt := "hi there" // var txt string = "hello there"
	var text string = "hello there"

	for i := 0; i < len(txt); i++{
		fmt.Println(txt[i]) // it gives the ASCII value of the character
		fmt.Printf("%c\n",txt[i]);
	}

	for idx, val := range text{
		fmt.Printf("index:%d value: %c\n",idx, val)
	}

	str := "हाय😊"
	for _, val := range str{
		fmt.Printf("%c", val)
		fmt.Println(val) // Break characters, Print garbage, Split emojis into pieces
	}

	r := []rune(txt)

	for i := 0; i < len(r); i++{
		fmt.Printf("%c\n", r[i])
	}
}