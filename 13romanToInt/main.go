package main

import "fmt"

func romanToInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if i+3 <= len(s) {
			if s[i:i+3] == "III" {
				n += 3
				i += 2
				continue
			} else if s[i:i+3] == "XXX" {
				n += 30
				i += 2
				continue
			} else if s[i:i+3] == "CCC" {
				n += 300
				i += 2
				continue
			} else if s[i:i+3] == "MMM" {
				n += 3000
				i += 2
				continue
			} else if s[i:i+3] == "VII" {
				n += 7
				i += 2
				continue
			} else if s[i:i+3] == "LXX" {
				n += 70
				i += 2
				continue
			} else if s[i:i+3] == "DCC" {
				n += 700
				i += 2
				continue
			}
		}
		if i+2 <= len(s) {
			if s[i:i+2] == "IV" {
				n += 4
				i++
				continue
			} else if s[i:i+2] == "XL" {
				n += 40
				i++
				continue
			} else if s[i:i+2] == "CD" {
				n += 400
				i++
				continue
			} else if s[i:i+2] == "VI" {
				n += 6
				i++
				continue
			} else if s[i:i+2] == "LX" {
				n += 60
				i++
				continue
			} else if s[i:i+2] == "CD" {
				n += 600
				i++
				continue
			} else if s[i:i+2] == "IX" {
				n += 9
				i++
				continue
			} else if s[i:i+2] == "XC" {
				n += 90
				i++
				continue
			} else if s[i:i+2] == "CM" {
				n += 900
				i++
				continue
			} else if s[i:i+2] == "II" {
				n += 2
				i++
				continue
			} else if s[i:i+2] == "XX" {
				n += 20
				i++
				continue
			} else if s[i:i+2] == "CC" {
				n += 200
				i++
				continue
			} else if s[i:i+2] == "MM" {
				n += 2000
				i++
				continue
			}
		}
		if s[i] == 'I' {
			n++
		} else if s[i] == 'V' {
			n += 5
		} else if s[i] == 'X' {
			n += 10
		} else if s[i] == 'L' {
			n += 50
		} else if s[i] == 'C' {
			n += 100
		} else if s[i] == 'D' {
			n += 500
		} else if s[i] == 'M' {
			n += 1000
		}
	}
	return n
}
func main() {
	roman := "XXI"
	fmt.Printf("%v\n %v", roman, romanToInt(roman))
}
