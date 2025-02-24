/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate different ways to declare Constants in go
ABOUT: This program declare and use Constant values
DATE: 29-01-2025
=============================================================================
*/

package main //package main is used when you're writing an executable program (not a library).

import "fmt"

func main() {

	// Create a constant
	const pi float64 = 3.1415
	const radis = 5
	const name string = "aswin"
	area := pi * radis * radis //Type of variable area is automatically found from the result of the output
	fmt.Println("Area is:", area)

	// Constants cannot be assigned another value once it is defines
	fmt.Println("Name is:", name)

}
