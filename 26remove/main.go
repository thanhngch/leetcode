package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
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
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	str = strings.Trim(str, " ")
	str = strings.Split(str, ".")[0]
	if len(str) == 0 {
		return 0
	}
	if str[0] == '+' {
		if len(str) > 1 && unicode.IsNumber(rune(str[1])) {
			str = str[1:]
		}
	}
	isNegative := false
	if str[0] == '-' {
		isNegative = true
		if len(str) > 1 && unicode.IsNumber(rune(str[1])) {
			str = str[1:]
		}
	}
	if !unicode.IsNumber(rune(str[0])) && str[0] != byte('-') {
		return 0
	}
	i, _ := strconv.Atoi(strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r) && r != '-'
	}))
	if i < math.MinInt32 {
		return math.MinInt32
	}
	if i > math.MaxInt32 {
		return math.MaxInt32
	}
	if isNegative {
		return -i
	}
	return i
}
func main() {
	i := myAtoi("123a")
	fmt.Println(i)
}
