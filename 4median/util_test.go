package main

import "testing"

func TestMin(t *testing.T) {
	cases := []struct {
		in1, in2, want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{5, 5, 5},
	}
	for _, c := range cases {
		got := min(c.in1, c.in2)
		if got != c.want {
			t.Errorf("min(%d, %d) == %d, want %d ", c.in1, c.in2, got, c.want)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		in1, in2, want int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{5, 5, 5},
	}
	for _, c := range cases {
		got := max(c.in1, c.in2)
		if got != c.want {
			t.Errorf("max(%d, %d) == %d, want %d ", c.in1, c.in2, got, c.want)
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		nums   []int
		middle float64
		want   int
	}{
		{[]int{2}, 1.5, 0},
		{[]int{2, 3}, 1.5, 0},
		{[]int{2, 3, 4}, 1.5, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1.5, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2.5, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 6.5, 6},
		{[]int{1, 5, 6, 8, 11, 12, 15}, 12.5, 6},
		{[]int{1, 5, 6, 8, 11, 12, 15}, 15.5, 7},
		{[]int{1, 5, 6, 8, 11, 12, 15}, 0.5, 0},

		{[]int{0, 2, 4, 6, 8, 10}, 5.5, 3},
		{[]int{1, 3, 5, 7, 9}, 5.5, 3},
		{[]int{1, 3, 5, 7, 9, 11, 13}, 5.5, 3},
		{[]int{1, 3, 5, 7, 9, 11}, 5.5, 3},

		{[]int{1, 2}, 1.5, 1},
		{[]int{1, 3, 5, 7}, 1.5, 1},
	}
	for _, c := range cases {
		got := divide(c.nums, c.middle)
		if got != c.want {
			t.Errorf("divide return %d, want %d ", got, c.want)
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{0, 2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}, 5.0},
		{[]int{0, 2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11, 13}, 6.0},
		{[]int{0, 2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}, 5.5},
		{[]int{1, 2}, []int{3}, 2.0},
		{[]int{2, 3}, []int{1}, 2.0},
		{[]int{2, 3}, []int{1, 5}, 2.5},
		{[]int{1, 2}, []int{1, 3, 5, 7}, 2.5},
		{[]int{2, 3}, []int{1, 5}, 2.5},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1}, []int{1}, 1.0},
		{[]int{2}, []int{3}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{100001}, []int{100000}, 100000.5},
		{[]int{5, 8}, []int{1, 2, 3, 4, 6, 7, 9}, 5.0},
		{[]int{3}, []int{1, 2, 4, 5, 6, 7, 8, 9}, 5.0},
	}
	for _, c := range cases {
		got := findMedianSortedArrays(c.nums1, c.nums2)
		if got != c.want {
			t.Errorf("%v and %v divide return %f, want %f ", c.nums1, c.nums2, got, c.want)
		}
	}
}
