/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate if else statement in go
ABOUT: This program creates if else conditioning
DATE: 05-02-2025

SYNTAX:

if condition {
	//code block
} else {
	//code block
}
=============================================================================
*/

package main

import "fmt"

func main() {

	a := 5
	if a > 5 {
		fmt.Println(a, "is greater than 5")
	} else {
		fmt.Println(a, "is less than 5")
	}
}
