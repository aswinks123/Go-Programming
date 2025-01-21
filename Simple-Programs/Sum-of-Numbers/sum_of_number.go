/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: Sum of 2 numbers
About: To add two numbers
DATE: 21-01-2025
=============================================================================
*/

// Every Go program starts with a package declaration.
package main

// Importing the fmt package to use the println function
import "fmt"

// Function to find the sum

func add(x int, y int) int {

	var output = x + y
	return output
}

// Program starts from here
func main() {
	fmt.Println("Program to find the sum of 2 numbers")

	num1 := 3
	num2 := 2

	result := add(num1, num2)
	fmt.Println(result)

}
