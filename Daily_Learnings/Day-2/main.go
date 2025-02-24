/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 24-02-2025

Problem: Given an integer slice nums, find the maximum number in the slice.
You must implement a function that returns the maximum number.

Example
Input: nums = [3, 1, 35, 7, 9, 5]
Output: 35
============================================================================
*/

package main

import "fmt"

//Function to find the largest number in the slice
func largenumber(nums []int) (int, error) { //function that returns int and an error

	if len(nums) == 0 { // checking if slice is empty or not. Return an error message if empty. 0 is returned because the function expects an int return value
		return 0, fmt.Errorf("Empty slice. Please add value to slice first")
	}

	large := nums[0] //assume first value as large
	for i := 1; i < len(nums); i++ {
		if nums[i] > large { // checking each element after the first value to the "large" variable
			large = nums[i] // if the value is large, then assign that value to the variable "large"
		}
	}
	return large, nil //return the large and nil ( nil is returned because the function expects an error return )
}

func main() {
	var nums = []int{3, 1, 35, 7, 9, 5}
	var larger, err = largenumber(nums) // calling the function
	if err != nil {                     //if variable err is not empty, then print the error that was returned
		fmt.Print(err)
		return // Exit the program
	}
	fmt.Println("Largest element is:", larger) // if not an empty slice then print the largest element

}
