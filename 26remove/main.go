package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
			i++
		} else if i == len(nums)-1 {
			break
		} else {
			i = findNextInt(nums, index, i)
		}
	}
	return index + 1
}

func findNextInt(nums []int, index int, i int) int {
	step := 1
	for {
		if i == len(nums)-1 {
			break
		}
		if i < len(nums)-1 && nums[i] == nums[index] && nums[i] != nums[i+1] {
			return i + 1
		}
		if i < len(nums) && nums[i] <= nums[index] {
			i += step
			step *= 2
		} else {
			i -= (step / 2) - 1
			step = 1
		}
	}
	return i
}

func main() {
	fmt.Println("Let's run test")
}
