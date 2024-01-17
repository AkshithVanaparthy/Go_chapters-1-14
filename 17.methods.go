package main

import (
	"fmt"
)

type Employe struct {
	name     string
	salary   int
	currency string
}

type Employeee struct {
	name string
	age  int
}

//Method with value receiver

func (e Employeee) changeName(newName string) {
	e.name = newName
}

//Method with pointer receiver

func (e *Employeee) changeAge(newAge int) {
	e.age = newAge
}

func (e Employe) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	emp1 := Employe{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	e := Employeee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("\nEmployee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("\nSum is", sum)
}
