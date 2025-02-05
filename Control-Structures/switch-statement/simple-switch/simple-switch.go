/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate Switch statement in go
ABOUT: This program creates Switch conditioning
DATE: 05-02-2025

SYNTAX:

switch expression {
case value1:
    // Code block if expression == value1
case value2:
    // Code block if expression == value2
default:
    // Code block if no cases match
}
=============================================================================
*/

package main

import "fmt"

func main() {

	day := "Friday"

	switch day {
	case "Thursday":
		fmt.Println("Start of work week!")
	case "Friday":
		fmt.Println("Weekend is coming!")
	default:
		fmt.Println("Regular Day!")
	}

}
