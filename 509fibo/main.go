package main

import "fmt"

func fib(N int) int {
	f := fibonacci()
	for i := 0; i < N; i++ {
		f()
	}
	return f()
}

func fibonacci() func() int {
	a := 1
	b := 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v\n", fib(i))
	// }
	fmt.Printf("%v\n", fib(35))
}
