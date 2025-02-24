/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 22-02-2025

Problem: Write a function in Go to check whether a given number is prime or not.
Input: n = 7
Output: true .Because  7 is only divisible by 1 and 7
Input: n = 10
Output: false
Explanation: 10 is not a prime number because it is divisible by 1, 2, 5, 10.
============================================================================
*/

package main

import (
	"fmt"
)

func checkprime(num int) bool { // function to check whether a numbr is prime or not. Return a boolean value
	if num == 1 { // if number is 1 the return false. 1 is not a prime number by default
		return false
	}

	for i := 2; i < num; i++ { //check the number from 2 to n-1 number
		if num%i == 0 { //if reminder of any number between 2 and n-1 is 0 , then that is not a prime number
			return false
		}
	}
	return true // return true means the number is prime. Because the other functions didn't return any value
}

func main() {

	var isprime bool
	num := 7
	isprime = checkprime(num) // Function call

	if isprime { // if isprime is true then below code executes
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}

}
