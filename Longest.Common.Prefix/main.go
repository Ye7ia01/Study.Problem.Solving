package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	// 1) If array only have one element, return that element as the common prefix.
	if len(strs) <= 1 {
		return strs[0]
	}

	// 2) Get the minimum len string
	minLength := 200
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}

	commonPrefix := ""
	breakSignal := false
	// 3) Iterate based on strings index till miLength
	for strIdx := 0; strIdx < minLength; strIdx++ {
		// Iterate on Array for each index
		for arrIdx := 0; arrIdx < len(strs)-1; arrIdx++ {
			// Compare characters at the current index of adjacent strings
			if strs[arrIdx][strIdx] != strs[arrIdx+1][strIdx] {
				breakSignal = true
				break
			}

			// Base condition, all are equal
			if arrIdx == len(strs)-2 {
				commonPrefix += string(strs[arrIdx][strIdx])
			}
		}

		if breakSignal {
			break
		}
	}
	return commonPrefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println("Longest Common Prefix:", result)
}
