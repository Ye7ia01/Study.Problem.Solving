package main

import "fmt"

/*
 Using memo the O(n) solution for climbing stairs problem.
 O(n) space complexity as well.
*/

/*
	BEST SOLUTION in terms of time and space complexity is the following
	func climbStairs(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    a, b := 1, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
	}
*/
func climbStairs(n int) int {
	memo := make(map[int]int)
	var helper func(int) int
	helper = func(n int) int {

		if n == 0 || n == 1 {
			return 1
		}
		// If cached FIB(n)
		if val, ok := memo[n]; ok {
			return val
		}
		// else calculate it and cache it
		memo[n] = helper(n-1) + helper(n-2)
		return memo[n]
	}
	return helper(n)
}
func main() {
	fmt.Println(climbStairs(44))
}
