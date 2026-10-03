package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// Newton's method Formula
	// x(n+1) = 0.5 * (x(n) + (x / x(n)))
	current := float64(x)
	for {
		next := 0.5 * (current + float64(x)/current)
		// stop iterating when INT(next value) = INT(current value)
		// i.e. when more iterations just increase accuracy and will not change the integer value of the result
		if int(next) == int(current) {
			break
		}
		current = next
	}
	return int(current)
}

func main() {

	fmt.Println(mySqrt(8))
}
