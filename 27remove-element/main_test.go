package main

import "testing"

var cases = []struct {
	in   []int
	ele  int
	want int
}{
	{[]int{}, 2, 0},
	{[]int{3}, 2, 1},
	{[]int{2, 3, 3, 2}, 2, 2},
	{[]int{2, 3, 3, 2}, 3, 2},
	{[]int{1, 3, 2, 3, 4, 3}, 3, 3},
	{[]int{4, 5}, 4, 1},
}

func TestRemoveElement(t *testing.T) {
	for _, c := range cases {
		got := removeElement(c.in, c.ele)
		if got != c.want {
			t.Errorf("removeElement %v, got %d, want %d ", c.in, got, c.want)
		}
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, c := range cases {
		got := removeElement2(c.in, c.ele)
		if got != c.want {
			t.Errorf("removeElement2 %v, got %d, want %d ", c.in, got, c.want)
		}
	}
}

func TestRemoveElement3(t *testing.T) {
	for _, c := range cases {
		got := removeElement3(c.in, c.ele)
		if got != c.want {
			t.Errorf("removeElement3 %v, got %d, want %d ", c.in, got, c.want)
		}
	}
}
