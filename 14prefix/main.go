package main

func longestCommonPrefix(strs []string) string {
	index := 0
	longPrefix := ""
	if len(strs) == 0 {
		return longPrefix
	}
	for {
		if len(strs[0]) <= index {
			return longPrefix
		}
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= index {
				return longPrefix
			}
			if strs[i][index] != strs[0][index] {
				return longPrefix
			}
		}
		longPrefix = longPrefix + string(strs[0][index])
		index++
	}
	return longPrefix
}

func main() {
	strs := []string{"thanh", "th"}
	longestCommonPrefix(strs)
}
