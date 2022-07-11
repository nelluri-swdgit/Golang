package main

import "fmt"

func main() {

	// short declaration operator (sdo)
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

	TypeCoversion()

}

func TypeCoversion(){
	// declare a variable with a built-in type
	var a int

	// create your own type and declare a  variable with that
	type chilli int
	var b chilli 

	a = 42
	b = 44

	// prints the values and their existing types
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// try dynamically changing the variable types and it will not work
	//  i.e. a = b will not work because both of them are declared as different types
	// and GoLang is a static programming language

	// now we try the type conversion, which is called as type casting in other programming languages
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
