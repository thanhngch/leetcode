package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMedianSortedArrays(nums2, nums1)
	}
	if len(nums2) == 0 {
		lenN1 := len(nums1)
		if lenN1%2 == 1 {
			return float64(nums1[lenN1/2])
		}
		return float64(nums1[(lenN1/2)-1]+nums1[(lenN1/2)]) / 2.0

	}
	return partion(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1)
}

func main() {
	m := findMedianSortedArrays([]int{3}, []int{1, 2, 4, 5, 6, 7, 8, 9})
	fmt.Println(m)
}
