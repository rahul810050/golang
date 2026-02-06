package main

import (
	"fmt"
)

func main(){
	fmt.Println("Slice in golang")
	// length vs capacity
	// length is usable elements in the slice 
	// capacity is allocated memory
	// [_ _ _ | _ _] the first three are length(how many elements is used) -- the last two is till 5 it can be filled(sort of allocated memory)
	nums := []int{1,2,3}
	fmt.Println(nums)

	nums = append(nums, 100)
	fmt.Println(nums)

	num2 := make([]int, 3, 5);
	fmt.Println(len(num2)) // 3
	fmt.Println(cap(num2)) // 5

	a := nums[1:3]
	fmt.Println(a)

	fruits := []string{}
	var n int 
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		var fru string
		fmt.Scan(&fru)
		fruits = append(fruits, fru)
	}

	for _, fru := range fruits{
		fmt.Println(fru)
	}


	// for i := 0; i < len(nums); i++{
	// 	fmt.Println(nums[i]);
	// }

	// for idx, val := range nums {
	// 	fmt.Println(idx, val)
	// }

	// for _, v := range nums{
	// 	fmt.Println(v);
	// }
}