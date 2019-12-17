package main

import "fmt"

func addBinary(a string, b string) string {
	result := ""
	remember := 0
	charToInt := map[byte]int{
		'1': 1,
		'0': 0,
	}
	i := len(a) - 1
	j := len(b) - 1
	for i > -1 || j > -1 {
		var ai byte = '0'
		if i > -1 {
			ai = a[i]
		}
		var bj byte = '0'
		if j > -1 {
			bj = b[j]
		}
		sum := charToInt[ai] + charToInt[bj] + remember
		switch sum {
		case 0:
			result = "0" + result
			remember = 0
			break
		case 1:
			result = "1" + result
			remember = 0
			break
		case 2:
			result = "0" + result
			remember = 1
			break
		case 3:
			result = "1" + result
			remember = 1
			break
		}
		i--
		j--
	}
	if remember == 0 {
		return result
	}
	return "1" + result
}
func main() {
	fmt.Printf("%v\n", addBinary("11", "100")) // 111
	fmt.Printf("%v\n", addBinary("11", "11"))  // 110
}
