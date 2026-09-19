package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Transform Array into a string representation */
func formatResult(ans []int) string {
	var builder strings.Builder
	for _, num := range ans {
		builder.WriteString(strconv.Itoa(num))
	}

	return builder.String()
}

/* Remove Leading Zeros in an array of integers */
func trimArray(ans []int) []int {
	leadingZeroes := 0
	for i := 0; ans[i] == 0 && i < len(ans); i++ {
		leadingZeroes++
	}
	return ans[leadingZeroes:]
}

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// Array holding answer
	// lenght will be max the addition on the lengths of the two numbers
	ans := make([]int, len(num1)+len(num2))

	// Loop Backwards
	for i := len(num1) - 1; i >= 0; i -= 1 {
		for j := len(num2) - 1; j >= 0; j -= 1 {
			// multiply and add to the current position in the answer array
			mull := int(num1[i]-'0')*int(num2[j]-'0') + ans[i+j+1]
			ans[i+j+1] = mull % 10 // single digit result (e.g. 15 -> 5)
			ans[i+j] += mull / 10  // carry and add on the next digit (e.g. 15 -> 1)
		}
	}
	return formatResult(trimArray(ans))
}

func main() {
	result := multiply("408", "5")
	fmt.Printf("result = %s\n", result)
}
