package main

import "fmt"

func reduce(nums []int, reducer func(int, int) int, init int) int {
	acc := init
	for i := 0; i < len(nums); i++ {
		acc = reducer(acc, nums[i])
	}
	return acc
}

func singleNumber(nums []int) int {
	return reduce(nums, func(acc int, value int) int {
		return acc ^ value
	}, 0)
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
