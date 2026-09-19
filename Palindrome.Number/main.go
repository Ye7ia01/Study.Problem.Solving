package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	reverse := 0 // int holding reverse value of original x
	temp := int(math.Abs(float64(x)))
	for temp != 0 {
		/*
			the following line combines 2 logics:
			1. temp % 10 grabs the current last digit of what's left of the number
			2. reverse * 10 shifts the digits already collected one place left,
			then + digit drops the new digit into the ones place
			(e.g. reverse=12, next digit=1 -> 12*10+1 = 121)
		*/
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}
	return reverse == int(math.Abs(float64(x)))
}

func main() {

	res := isPalindrome(121)
	if res {
		fmt.Printf("Is Palindrom")
	} else {
		fmt.Printf("Not Palindrome")
	}
}
