package main
import "fmt"


func PrintInNumerals(){
	a := 444

	fmt.Printf("%T\n", a)
	fmt.Printf("value of the variable in decimal, binary, hexadecimal is "+ "%d\t%b\t%#x", a, a, a)
}

func DiffOperatorVars(){
	g := (1== 1)
	h := (1<= 2)
	i := (1>= 0)
	j := (1!= 2)
	k := (1 < 3)
	l := (1 > 0)

	fmt.Println("Variables created using operators have corresponding values of type bool", g, h, i, j, k, l)
}


func Constants(){
	 const(
		b = 44
		c int8 = 4
	 )
	 fmt.Println("value of the untyped constsnt b is:", b)
	 fmt.Printf("Type of the untyped constant b is:" + "%T\n", b)
	 fmt.Println(" ")
	 fmt.Println("value of the untyped constsnt c is:", c)
	 fmt.Printf("Type of the typed constant c is:" + "%T\n", c)
}


func CrudeBitShifting(){
	a := 444

	fmt.Printf("%T\n", a)
	fmt.Printf("value of the variable in decimal, binary, hexadecimal is "+ "%d\t%b\t%#x\n", a, a, a)

	d := a << 1
	fmt.Printf("%T\n", d)
	fmt.Printf("value of the variable in decimal, binary, hexadecimal is "+ "%d\t%b\t%#x", d, d, d)


}

func RawString(){
	s := `I am a Raw String Literal assigned to an Identifer S. 
you may find it here....!'`
	fmt.Println(s)
}

func iotaConstants(){
	type year int
	const(
		_ = iota
		next year = 2022 + iota
		next1 = next + iota
		next2 = next1 + iota
		next3 = next2 + iota
	)

	fmt.Println(next, next+1, next+2, next+3)
}

// main function that excecutes the functions and deals with main tasks of the script
func main(){
	fmt.Println("Here on i will execute and print the functions written in this file..!")
	fmt.Println(" ")

	fmt.Println("Now executing function: PrintInNumerals ...........")
	PrintInNumerals()
	fmt.Println(" ")

	fmt.Println("Now executing function: DiffOperatorVars ...........")
	DiffOperatorVars()
	fmt.Println(" ")

	fmt.Println("Now executing function: Constants ...........")
	Constants()
	fmt.Println(" ")

	fmt.Println("Now executing function: CrudeBitShifting ...........")
	CrudeBitShifting()
	fmt.Println(" ")

	fmt.Println("Now executing function: RawString ...........")
	RawString()
	fmt.Println(" ")

	fmt.Println("Now executing function: iotaConstants ...........")
	iotaConstants()
	fmt.Println(" ")

}