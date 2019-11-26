package main

import "fmt"

func reverse(x int) int {
	n := 0
	for x != 0 {
		n *= 10
		n += x % 10
		x = x / 10
	}
	if n > 2147483647 || n < -2147483648 {
		return 0
	}
	return n
}

func main() {
	x := reverse(-123456789)
	fmt.Printf("%d\n", x)
}
