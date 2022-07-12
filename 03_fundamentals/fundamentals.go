package  main

/* multi import possibility
	import {
		"fmt"
		"runtime"
	}
*/
import "fmt"
import "runtime"

// varibale declaration with the zero values
var x bool
var y int8
var z float64
var r uint8
var s int
var n uint


// checking runtime package - default from Golang
func RunTime(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func NumericStuff(){
	// Simple code to understand the available Numeric types in Golangg
	fmt.Printf("%T" + " : takes values only between -128 to 127 \n", y)
	y = -128
	fmt.Println(y)

	fmt.Printf("%T" + " : Only way to adeclare with fractional values - 64-bit \n", z)
	z = 222.69
	fmt.Println(z)

	fmt.Printf("%T" + " : takes values only between 0 to 255 \n", r)
	r = 255
	fmt.Println(r)

	fmt.Printf("%T" + " : takes values 32 or 64 bits8 - depends on machine \n", s)
	s = 2555555555
	fmt.Println(s)

	fmt.Printf("%T" + " : takes values 32 or 64 bits - depends on machine \n", n)
	n = 9999999999
	fmt.Println(n)
}


func Stringstuff(){
	st := "Hello, Ramu"
	fmt.Println(st)
	fmt.Printf("%T\n", st)

	// Slicing the string into bytes and checking the type of resulting data
	bst := []byte(st)
	fmt.Println(bst)
	fmt.Printf("%T\n", bst)

	// Looping over a string using its length as range
	for i := 0; i < len(st); i++ {
		fmt.Printf("%#U", st[i])
	}

	fmt.Println(" ")

	// Looping over a string and assigning the index and value to variables
	for p, v := range st {
		fmt.Printf("for the index %d we have the hex value %#x\n", p, v)

	}
	

}

/*  checking with constants and iota
	const declares the variable types, 
	whilst iota by default declares int type to the variables in an incremental manner (increments by 1)
*/
const (
	a = iota
	b
	c
)
func ReviewConstants(){

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

// bit shifting strategy using crude vs iota approach

func basicbitshift(){
	f := 1
	fmt.Printf("value of f before basic bit shifting " + "%d\t\t\t%b\n", f, f)
	g := f << 3
	fmt.Printf("value of g after basic bit shifting " +"%d\t\t\t%b\n", g, g)
}

func iotabitshifting(){
	const (
		_ = iota
		kb = 1 << (iota*10)
		mb = 1 << (iota*10)
		gb = 1 << (iota*10)
		tb = 1 << (iota*10)
	)

	fmt.Printf("value of kb with iota bit shifting " + "%d\t\t\t%b\n", kb, kb)
	fmt.Printf("value of mb with iota bit shifting " + "%d\t\t%b\n", mb, mb)
	fmt.Printf("value of gb with iota bit shifting " + "%d\t\t%b\n", gb, gb)
	fmt.Printf("value of tb with iota bit shifting " + "%d\t%b\n", tb, tb)
}


// Main Function Section, that can execute all the fuctions in the script body
func main(){
	fmt.Println(x)
	y := false
	fmt.Println(x==y)

	RunTime()
	fmt.Println()

	NumericStuff()
	fmt.Println()

	Stringstuff()
	fmt.Println()
	
	ReviewConstants()
	fmt.Println()
	
	basicbitshift()
	fmt.Println()
	
	iotabitshifting()
}