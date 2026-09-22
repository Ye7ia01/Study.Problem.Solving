package main

import (
	"fmt"
	"sort"
)

// returns index of searched element O(log n)
func binarySearch(target int, arr [][2]int) int {
	var low int = 0
	var high int = len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2 // Low + is necessary since (high-low)/2 gives you the midle between
		// both low and high and not the index mid in the array so you need to add Low again
		// so that the mid is correctly positioned, this is visible when we try to get the middle when low ! 0
		// e.g try to get the mid for low=3 and Hiht = 6, only (high-Low)/2 = 2
		// but the actual mid index should be 3 (current index) + 2 (medium index) = 5
		if arr[mid][0] == target {
			return arr[mid][1] // return index
		} else if target > arr[mid][0] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Overall time complexity O(n*Log n)
// O(n) for creating the copy + O(nlogn) for sorting + O(nlogn) for binary search in the loop
func twoSum(numbers []int, target int) []int {
	// Create a copy of the array with indices O(n)
	numbersWithIndices := make([][2]int, len(numbers))
	for i, num := range numbers {
		numbersWithIndices[i] = [2]int{num, i}
	}

	// Sort the array based on values O(nlogn)
	sort.Slice(numbersWithIndices, func(i, j int) bool {
		return numbersWithIndices[i][0] < numbersWithIndices[j][0]
	})

	for i := 0; i < len(numbersWithIndices); i++ {
		complement := target - numbersWithIndices[i][0]
		scndIndx := binarySearch(complement, numbersWithIndices[i+1:])
		if scndIndx != -1 {
			return []int{numbersWithIndices[i][1], scndIndx}
		}
	}
	return []int{0, 0}
}

func main() {

	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
