/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate Switch Statement
About: This program checkes when is saturday comparing to today's date
DATE: 21-01-2025
=============================================================================
*/

// Every Go program starts with a package declaration.
package main

// Importing the fmt, time package to use the println and time related function
import (
	"fmt"
	"time"
)

// Program starts from here
func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()

	fmt.Println("Today is : ", today)
	fmt.Print("Saturday is: ")

	// Switch statement to check the condition
	switch time.Saturday {

	// today stores todays day.
	case today + 0:
		fmt.Println("Today.")

	// today+1 stores tomorrow's day.
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In Three days.")
	default:
		fmt.Println("Too far away.")
	}
}
