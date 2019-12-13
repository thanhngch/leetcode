package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	return search(nums, target, 0, len(nums))
}

func search(nums []int, target int, start int, end int) int {
	if start == end {
		return start
	}
	middleIndex := (start + end) / 2
	if nums[middleIndex] == target {
		return middleIndex
	} else if nums[middleIndex] < target {
		return search(nums, target, middleIndex+1, end)
	} else if nums[middleIndex] > target {
		return search(nums, target, start, middleIndex)
	}
	return 0
}

func main() {
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 0)) // 0
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 4)) // 2
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 6)) // 3
}
