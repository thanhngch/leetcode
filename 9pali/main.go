package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	n := 0
	oldX := x
	for x != 0 {
		n *= 10
		n += x % 10
		x = x / 10
	}
	if oldX == n {
		return true
	}
	return false
}
func main() {
	x := isPalindrome(0)
	fmt.Printf("%v\n", x)
}
