package main

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 5, 5, 5, 5, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 1}, []int{1}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1}},
		{[]int{1}, []int{1}},
	}
	for _, c := range cases {
		gotLen := removeDuplicates(c.in)
		for i := 0; i < len(c.want); i++ {
			if c.in[i] != c.want[i] {
				t.Errorf("removeDuplicates(%v), want %v ", c.in, c.want)
				break
			}
		}
		if gotLen != len(c.want) {
			t.Errorf("removeDuplicates(%v) len %d, want %d ", c.in, gotLen, len(c.want))
		}
	}
}
