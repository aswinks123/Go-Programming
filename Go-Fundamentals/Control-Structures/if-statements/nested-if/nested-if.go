/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate nested if statement in go
ABOUT: This program creates nested if conditioning
DATE: 05-02-2025

=============================================================================
*/

package main

import (
	"fmt"
)

func main() {

	age := 22

	if age > 20 {
		fmt.Println("Adult")

		if age >= 19 {
			fmt.Println("Can drink Alcohol")
		} else {
			fmt.Println("Cannot drink Alcohol !")
		}
	} else {
		fmt.Println("Minor!!")
	}

}
