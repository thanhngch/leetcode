package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"babad", "bab"},
		{"ccc", "ccc"},
		{"abbba", "abbba"},
		{"abbcs", "bb"},
	}
	for _, c := range cases {
		got := longestPalindrome(c.in)
		if got != c.want {
			t.Errorf("longest(%s) == %s, want %s ", c.in, got, c.want)
		}
	}
}
