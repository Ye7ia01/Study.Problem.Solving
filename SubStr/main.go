package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	ret := -1
	for i := 0; i < len(haystack); i++ {
		if needle[0] != haystack[i] {
			continue
		}
		k := i
		for j := 0; j < len(needle); j++ {
			if k > len(haystack)-1 {
				break
			}
			if needle[j] != haystack[k] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
			k++
		}

	}
	return ret
}

func main() {
	fmt.Println(strStr("mississippi", "sipp"))
}
