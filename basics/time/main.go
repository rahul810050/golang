package main

import (
	"fmt"
	"time"
)


func main(){
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Printf("type of currTime %T\n", currTime) // golang has their own time data type -output-> time.Time

	formattedTime := currTime.Format("02-01-2006, 15:04:05")
	fmt.Println(formattedTime)

	layout_str := "02-01-2006"
	dateStr := "08-02-2026"

	formated_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("formatted time:", formated_time)


	// add 1 more day in currentTime
	new_date := currTime.Add(24 * time.Hour)
	fmt.Println(new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println(formatted_new_date)
}