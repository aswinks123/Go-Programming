/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate Switch fallthrough in go
ABOUT: This program creates Switch fallthrough in go
DATE: 05-02-2025

By default, Go does not automatically execute the next case after a match.
However, you can use fallthrough to force execution of the next case.(Even if the condition case fails)
============================================================================
*/

package main

import "fmt"

func main() {
	number := 1

	switch number {

	case 0:
		fmt.Println("Not same")
	case 1:
		fmt.Println("Match")
		fallthrough // This forces go to execute the next statement. It ignores the condition checking
	case 2:
		fmt.Println("This runs eventhough it is not a match because of the fallthrough statement in previous case")
	}

}
