package main

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{123, 321},
		{1, 1},
		{-1, -1},
		{1000, 1},
		{-1000, -1},
		{-12345678, -87654321},
		{1534236469, 0},
	}
	for _, c := range cases {
		got := reverse(c.in)
		if got != c.want {
			t.Errorf("reverse(%d) == %d, want %d ", c.in, got, c.want)
		}
	}
}
