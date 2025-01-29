/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate different ways to declare variables in go
ABOUT: This program declare and assign varibale values
DATE: 29-01-2025
=============================================================================
*/

package main //package main is used when you're writing an executable program (not a library).

import "fmt"

func main() {
	// Declaring and assigning varibale separately
	var name string
	name = "Aswin"
	fmt.Println("Hello", name)

	// Decalring and assigning in same line
	var age int = 25
	fmt.Println("Your age is:", age)

	//Short variable decleration method
	country := "Canada"
	fmt.Println("Your country is:", country)

	//Multiple varibale declerations
	var provience, city string = "ON", "Toronto"
	fmt.Printf("Your Provience is %v and City is %v \n", provience, city)

	//Zero Values (Uninitialized variables in Go are automatically assigned a "zero value)
	//int = 0, float64 = 0.0, bool = false, string = ""
	var isavailable bool
	fmt.Println(isavailable)

}
