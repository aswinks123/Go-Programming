/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 22-02-2025

Problem: Write a Go function to reverse a given array in place (without using an extra array).
Input: nums = [1, 2, 3, 4, 5]
Output: [5, 4, 3, 2, 1]
============================================================================
*/

package main

import "fmt"

func reversearray(nums []int) []int {

	for i := 0; i < len(nums)/2; i++ {
		x := nums[i]
		nums[i] = nums[len(nums)-(i+1)]
		nums[len(nums)-(i+1)] = x

	}
	return nums
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	reverse := reversearray(nums)
	fmt.Print("Reverse is ", reverse)

}
