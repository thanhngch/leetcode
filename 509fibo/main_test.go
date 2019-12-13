package main

import "testing"

func TestFibonaci(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{35, 9227465},
	}
	for _, c := range cases {
		got := fib(c.in)
		if got != c.want {
			t.Errorf("fibonaci(%v), want %v, got %v ", c.in, c.want, got)
			break
		}
	}
}
