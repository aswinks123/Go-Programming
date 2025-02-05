/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate if with short variable decleration statement in go
ABOUT: This program creates if short veriable decleration conditioning
DATE: 05-02-2025

=============================================================================
*/

package main

import "fmt"

func main() {
	if num := 5; num > 40 { // Declaring and defining a variable directly in loop. this variable is scoped to if else block only
		fmt.Println("Number is greater than 40")
	} else {
		fmt.Println("Number is less than 40")
	}

}
