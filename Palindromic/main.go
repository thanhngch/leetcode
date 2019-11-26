package main

import "fmt"

func longestPalindrome(s string) string {
	maxString := ""
	for i := range s {
		j, k := 0, 0
		for ; j+i < len(s) && i-j > -1; j++ {
			if s[i+j] != s[i-j] {
				break
			}
		}
		for ; k+i+1 < len(s) && i-k > -1; k++ {
			if s[i+k+1] != s[i-k] {
				break
			}
		}
		j--
		if 2*j+1 > len(maxString) {
			maxString = s[i-j : i+j+1]
		}
		if 2*k > len(maxString) {
			maxString = s[i-k+1 : i+k+1]
		}
	}
	return maxString
}
func main() {
	maxString := longestPalindrome("babad")
	fmt.Printf("%s\n", maxString)
}
