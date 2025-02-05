/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate Multiple case Switch statement in go
ABOUT: This program creates Multiple case Switch conditioning
DATE: 05-02-2025
============================================================================
*/

package main

import "fmt"

func main() {
	grade := "A"
	switch grade {
	case "A", "B": // Checks multiple cases in one statement. It run if any is True
		fmt.Println("Excellent")
	case "C", "D":
		fmt.Println("Good")
	default:
		fmt.Println("Need Improvement")

	}

}
