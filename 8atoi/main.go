package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	str = strings.Split(str, ".")[0]

	chars := []rune(str)

	if len(chars) == 0 {
		return 0
	}

	var result int
	isNegative := 1
	if chars[0] == '-' {
		isNegative = -1
		chars = chars[1:]
	} else if chars[0] == '+' {
		chars = chars[1:]
	} else {
		intChar := int(chars[0]) - int('0')
		if intChar < 0 || intChar > 9 {
			return 0
		}
	}

	for _, char := range chars {
		intChar := int(char) - int('0')
		if intChar >= 0 && intChar < 10 {
			result = result*10 + intChar
		} else {
			break
		}
		if isNegative == -1 {
			resultNe := result * isNegative
			if resultNe < math.MinInt32 {
				return math.MinInt32
			}
			if resultNe > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			if result < math.MinInt32 {
				return math.MinInt32
			}
			if result > math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}
	result = result * isNegative

	return result
}
func main() {
	i := myAtoi("123a")
	fmt.Println(i)
}
