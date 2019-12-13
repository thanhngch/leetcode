package main

import "fmt"

func plusOne(digits []int) []int {
	remember := false
	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] == 9 {
			remember = true
			digits[i] = 0
		} else {
			remember = false
			digits[i]++
		}
		if !remember {
			return digits
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Printf("%v\n", plusOne([]int{1, 2, 3}))    // 1,2,4
	fmt.Printf("%v\n", plusOne([]int{4, 3, 2, 1})) // 4,3,2,2
	fmt.Printf("%v\n", plusOne([]int{8, 9}))       // 9,0
	fmt.Printf("%v\n", plusOne([]int{8, 9, 9}))    // 9,0,0
	fmt.Printf("%v\n", plusOne([]int{9, 9, 9}))    // 1,0,0,0
}
