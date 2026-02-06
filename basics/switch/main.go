package main

import (
	"fmt"
)

func main(){
	fmt.Println("learning switch in golang")

	day := 5

	switch day {
	case 1:
		fmt.Println("it is day 1")
	case 2:
		fmt.Println("it is day 2")
	case 3:
		fmt.Println("it is day 3")
	default:
		fmt.Println("unknown day")
	}

	month := "jan"

	switch month{
	case "jan", "feb", "dec":
		fmt.Println("it is winter")
	case "apr", "may", "june":
		fmt.Println("it is spring")
	default:
		fmt.Println("season unknown")
	}

	tem := 25

	switch {
	case tem < 0:
		fmt.Println("it is freezing")
	case tem >= 0 && tem <= 10:
		fmt.Println("it is cold")
	case tem >= 10 && tem < 20:
		fmt.Println("it is cool")
	default:
		fmt.Println("it is hot")
	}
}