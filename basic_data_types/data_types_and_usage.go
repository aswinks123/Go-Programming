/*
=============================================================================
DEVELOPER: Aswin KS
PURPOSE: To demonstrate different date types in go
About: This program creates variables of different data types
DATE: 29-01-2025
=============================================================================
*/

package main //package main is used when you're writing an executable program (not a library).
import "fmt"

func main() {

	// Integers can be int, int8, int16,int32, and int64. It can be signed and unsigned(Only positive)
	var a int = -5
	var b uint = 20
	var c int8 = 126
	fmt.Println(a, b, c)

	// Floats can be float32 or float64
	var f1 float32 = 4.5
	var f2 float64 = 7.7747774536456456
	fmt.Println(f1, f2)

	//Basic Arithmetic Opeartions
	x := 5
	y := 10
	fmt.Println(x + y)

	//Boolean Operations
	isstudent := true
	isteacher := false

	fmt.Println(isstudent)
	fmt.Println(isteacher)

	//String Operations
	// Enclosed in "" , Default is empty string , can concatinate using '+' operator
	s1 := "Aswin"
	s2 := "KS"
	fmt.Println("Full name is:", s1+" "+s2)

	//Multiline String
	multistring := `This is
	a multi line
	string `
	fmt.Println(multistring)

	//Type Conversion : COnvert data from one type to another

	var value1 int = 10
	fmt.Println(value1)
	var value2 float32 = float32(value1)
	fmt.Printf("Converted float value is %.2f", value2)
}
