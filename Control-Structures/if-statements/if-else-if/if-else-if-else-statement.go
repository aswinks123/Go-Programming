/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate if else if else statement in go
ABOUT: This program creates if else if else conditioning
DATE: 05-02-2025

SYNTAX:

if condition {
	//code block
} else if condition {
	//code block
} else {
	//code block
}
=============================================================================
*/

package main

import "fmt"

func main() {

	a := -3
	if a > 0 {
		fmt.Println(a, "is positive value")
	} else if a < 0 {
		fmt.Println(a, "is negative value")
	} else {
		fmt.Println(a, "is Zero")
	}
}
