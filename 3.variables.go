package main

import (
	"fmt"
	"math"
)

func main() {
	// declaring a variable

	var age int
	println("My age is : ", age)

	//Q2
	// declaring a new variable
	var age2 int
	println("My age is: ", age2)
	age2 = 21 // assigment
	println("My age is: ", age2)
	age2 = 30
	println("My new age is: ", age2)

	//Q3
	var width, height int = 101, 60

	println("width is", width, "height is", height)
	var (
		name     = "Akshith"
		myage    = 19
		myheight int
	)
	println("my name is", name, ", age is", myage, "and height is", myheight)

	a, b := 10, 30 // declare variables a and b
	println("a is", a, "b is", b)
	b, c := 40, 60 // b is already declared but c is new
	println("b is", b, "c is", c)
	b, c = 80, 70 // assign new values to already declared variables b and c
	println("changed b is", b, "c is", c)
	a1, b1 := 17.8, 23.8
	c1 := math.Min(a1, b1)
	fmt.Println("Minimum value is", c1)
}
