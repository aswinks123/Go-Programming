/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 22-02-2025

Problem: Given an array of integers nums and an integer target, return all the indices of the two numbers that add up to target.
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: nums[0] + nums[1] = 2 + 7 = 9
============================================================================
*/

package main

import "fmt"

// twosum function to check the sum of values
func twosum(nums []int, target int) [][]int {
	result := [][]int{} // Slice to store all valid index pairs

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j}) // Add pair as a slice
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 5, 3, 8, 4} // Given Array
	target := 9
	result := twosum(nums, target) //Calling the userdefined function named twosum()
	fmt.Println(result)
}
