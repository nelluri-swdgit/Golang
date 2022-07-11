package main
import "fmt"

// declare variables with the types and they are called as zero values
var x int
var y string
var z bool

type pepper int
var p pepper
var b int

func main(){
	x := 42
	y := "James Bond 007"
	z := true

	//single var print statements
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//multi var print statement
	fmt.Println(x, y, z)

	// print all the variable values into a single string using fmt.Sprintf
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)

	// assigning a custom type to var p and printing
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	p = 444
	fmt.Println(p)

	// Type conversion
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	b = int(p)
	fmt.Println(b)
}

// for the details on the questions and solutions written - 
// refer the Ninja Level 1 section from excercises on Udemy