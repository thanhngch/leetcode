package main

import "testing"

func TestAtoI(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"123a", 123},
		{"-123a", -123},
		{"4193 with words", 4193},
		{"-91283472332", -2147483648},
		{"91283472332", 2147483647},
		{"words and 987", 0},
		{"+1", 1},
		{"   -42", -42},
		{"+-2", 0},
		{"9223372036854775807", 2147483647},
		{"9223372036854775809", 2147483647},
		{"9223372036854775808", 2147483647},
	}
	for _, c := range cases {
		got := myAtoi(c.in)
		if got != c.want {
			t.Errorf("myAtoi(%v) == %d, want %d ", c.in, got, c.want)
		}
	}
}
