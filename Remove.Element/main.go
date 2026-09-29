package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for x := 0; x < len(nums); x++ {
		if nums[x] != val {
			nums[i] = nums[x]
			i++
		}

	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	removeElement(nums, val)
	fmt.Printf("%v", nums)
}
