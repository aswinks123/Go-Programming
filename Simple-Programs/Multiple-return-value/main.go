/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate function that return multple value
About: This program creates a function that returns multiple values
DATE: 21-01-2025
=============================================================================
*/

// Every Go program starts with a package declaration.
package main

// Importing the fmt package to use the println function
import "fmt"

// Function to find the sum and difference
// x and y are call by value variables and sum and diff are return values
func calc(x int, y int) (sum, diff int) {
	sum = x + y
	diff = x - y
	// Return all the values specified on the function head
	return
}

// Program starts from here
func main() {
	fmt.Println("Program to find the sum and difference of 2 numbers")
	num1 := 3
	num2 := 2
	// Variables to store the values of sum and diff that was returned by the calc function
	sum, difference := calc(num1, num2)
	// Print the outputs line after another.
	fmt.Println("Sum of Numbers:", sum)
	fmt.Println("Difference of Numbers:", difference)

}
