package main

import "fmt"

func main() {
	//Different Type of Constant

	//Declaring a constant
	const (
		name    = "Akshith"
		age     = 19
		country = "India"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)
	//Boolean Constant

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst
	println(defaultBool)
	println(customBool)

	//Numeric Constant
	const a = 7
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
