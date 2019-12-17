package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i2 := 0
	arr := []int{}
	iArr := 0
	if n == 0 {
		return
	}
	for i1 := 0; i1 < m+n; i1++ {
		if iArr == len(arr) && nums1[i1] <= nums2[i2] && i1 < m {
			continue
		} else {
			if i1 < m {
				arr = append(arr, nums1[i1])
			}
			if i2 >= n || iArr < len(arr) && i2 < n && arr[iArr] < nums2[i2] {
				nums1[i1] = arr[iArr]
				iArr++
			} else {
				nums1[i1] = nums2[i2]
				i2++
			}
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%v\n", nums1) // [ 1, 2, 2, 3, 5, 6 ]

	nums1 = []int{1}
	merge(nums1, 1, []int{}, 0)
	fmt.Printf("%v\n", nums1) // [ 1 ]

	nums1 = []int{2, 0}
	merge(nums1, 1, []int{1}, 1)
	fmt.Printf("%v\n", nums1) // [ 1, 2 ]
}
