package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNumsPrev := 0
	maxNumsNext := nums[0]
	maxNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		maxNumsPrev = maxNumsNext
		maxNumsNext = max(nums[i], maxNumsPrev+nums[i])
		maxNumber = max(maxNumber, maxNumsNext)
	}
	return maxNumber
}

func main() {
	fmt.Printf("%v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
}
