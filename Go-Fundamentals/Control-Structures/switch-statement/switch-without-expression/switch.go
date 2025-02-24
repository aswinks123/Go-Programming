/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate Switch statement in go
ABOUT: This program creates Switch conditioning
DATE: 05-02-2025

If no expression is provided, switch behaves like an if-else structure.
============================================================================
*/

package main

import "fmt"

func main() {

	num := 10
	switch { // if no condition is provided switch act similar to if else statement
	case num < 0:
		fmt.Println("Negative number")
	case num > 0:
		fmt.Println("Positive number")
	case num == 0:
		fmt.Println("Zero")

	}

}
