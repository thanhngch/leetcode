package main

import "fmt"

func lengthOfLastWord(s string) int {
	end := len(s)
	i := len(s) - 1
	for ; i > -1; i-- {
		if s[i] == ' ' {
			end = i
		} else {
			break
		}
	}
	for ; i > -1; i-- {
		if s[i] == ' ' {
			break
		}
	}
	return end - i - 1
}

func main() {
	fmt.Printf("%v\n", lengthOfLastWord("Hello World")) // 5
	fmt.Printf("%v\n", lengthOfLastWord("Hello "))      // 5
	fmt.Printf("%v\n", lengthOfLastWord("  Hello  "))   // 5
	fmt.Printf("%v\n", lengthOfLastWord(" "))           // 0
	fmt.Printf("%v\n", lengthOfLastWord(""))            // 0
}
