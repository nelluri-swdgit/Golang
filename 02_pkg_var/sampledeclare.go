package main

import "fmt"

func main() {

	// short declaration oprator (sdo)
	// Declare a variale and assign a value to it
	// useful to declare a variable inside a function
	x := 10
	fmt.Println(x)
	y := 20
	fmt.Println(y)
	z := "Bond Bond..." + "James Bond"
	fmt.Println(z)

	// excluding sdo with the keyword var
	// use this when you need to declare a variable outside the function
	// suggestion: Limit the scope of the variable to the requird functions only
	var n = "Mercedes"
	fmt.Println(n)

	// Finding out the type of a variable
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// Primitive data types aka basic/built-in ex: bool, int, string
	// composite data types, which hold the data of other types ex: arrays, slice, structs

}
