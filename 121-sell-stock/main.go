package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minValue := prices[0]
	diff := 0
	for i := 0; i < len(prices); i++ {
		minValue = min(prices[i], minValue)
		diff = max(prices[i]-minValue, diff)
	}
	return diff
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
