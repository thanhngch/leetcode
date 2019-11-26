package main

import (
	"fmt"
	"math"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	lenSubstr := 2*numRows - 2
	resultS := ""
	maxI := int(math.Ceil(float64(len(s)) / float64(lenSubstr)))
	for i := 0; i < maxI; i++ {
		subString := s[i*lenSubstr : min((i+1)*lenSubstr, len(s))]
		resultS += string(subString[0])
	}

	for k := 1; k < lenSubstr/2; k++ {
		for i := 0; i < maxI; i++ {
			subString := s[i*lenSubstr : min((i+1)*lenSubstr, len(s))]
			if k < len(subString) {
				resultS += string(subString[k])
			}
			if lenSubstr-k < len(subString) {
				resultS += string(subString[lenSubstr-k])
			}
		}
	}
	for i := 0; i < maxI; i++ {
		subString := s[i*lenSubstr : min((i+1)*lenSubstr, len(s))]
		if lenSubstr/2 < len(subString) {
			resultS += string(subString[lenSubstr/2])
		}
	}
	return resultS
}

// Reduce is functional function
func Reduce(it []int, red func(int, int) int,
	accumulator int) int {
	for _, value := range it {
		accumulator = red(accumulator, value)
	}
	return accumulator
}

// Missing Number
func addTotal(nums []int) int {
	return Reduce(nums, func(a int, b int) int {
		return a + b
	}, 0)
}
func missingNumber(nums []int) int {
	n := len(nums)
	sum := addTotal(nums)
	return n*(n+1)/2 - sum
}
func missingNumber2(nums []int) int {
	missing := len(nums)
	for i, v := range nums {
		missing ^= i ^ v
	}
	return missing
}
func main() {
	s := convert("A", 1)
	fmt.Printf("%s\n", s)

	fmt.Printf("Missing %d", missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
