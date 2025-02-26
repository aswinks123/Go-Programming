/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 22-02-2025

Problem: Write a Go program to find the second largest element in an array without using built-in functions.
Input: nums = [10, 20, 4, 45, 99]
Output: 45
============================================================================
*/

package main

import "fmt"

func main() {

	nums := []int{10, 20, 4, 45, 99, 100}

	// Find the largest

	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}

	}
	// Now Find the second largest

	secondLargest := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > secondLargest && nums[i] < largest {
			secondLargest = nums[i]
		}
	}
	fmt.Println("Second largest is :", secondLargest)

}
