package main

import "fmt"

func removeElement(nums []int, val int) int {
	var begin = 0
	var end = len(nums) - 1
	if end == -1 {
		return 0
	}
	for end != 0 {
		for nums[begin] != val && begin < end {
			begin++
		}
		for nums[end] == val && begin < end {
			end--
		}
		if begin >= end {
			break
		}
		// swap
		var temp = nums[begin]
		nums[begin] = nums[end]
		nums[end] = temp

		begin++
		end--
	}
	if nums[begin] == val {
		return begin
	}
	return begin + 1
}

func removeElement2(nums []int, val int) int {
	var i = 0
	var n = len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}

func removeElement3(nums []int, val int) int {
	var i = 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	var nums = []int{2, 3, 4, 2}
	var length = removeElement3(nums, 2)
	fmt.Println(length)
}
