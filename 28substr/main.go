package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("chicken", "ken"))
	fmt.Println(strStr("chicken", "dmr"))
	fmt.Println(strStr("hello", "ll"))
}
