/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate boolean expressons in go
ABOUT: This program creates demonstrate boolean expressons in go
DATE: 05-02-2025

Go supports boolean expressions using:

Logical Operators: && (AND), || (OR), ! (NOT)
Comparison Operators: ==, !=, <, >, <=, >=

=============================================================================
*/

package main

import "fmt"

func main() {
	x, y := 5, 10

	if x > 0 && y > 0 {
		fmt.Println("Both x and y are positive")
	}
	if x < 0 && y < 0 {
		fmt.Println("Both x and y are Negative")
	}
	if x == 0 && y == 0 {
		fmt.Println("Both x and y are Zero")
	}

}
