/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 22-02-2025

Problem: Write a Go function to reverse a given array in place (without using an extra array).
Input: nums = [1, 2, 3, 4, 5]
Output: [5, 4, 3, 2, 1]
Note: This program uses 2 pointer approach
============================================================================
*/

package main

import "fmt"

func reversearray(nums []int) []int { // Function body

	left, right := 0, len(nums)-1 // set left as 0 and right as (length of array-1)

	for left < right { // Logic to swap each element from first to last position
		nums[left], nums[right] = nums[right], nums[left] // Swapping
		left++                                            // incrementing the left
		right--                                           // Decrementing the right
	}
	return nums // Return the array to called function

}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 1} // Array

	reverse := reversearray(nums)     // Calling the reversearray function
	fmt.Print("Reverse is ", reverse) // Print the result array

}
