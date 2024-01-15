package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	//Q2 Signed integer
	var a1 int = 70
	b1 := 105
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b1, unsafe.Sizeof(b1)) //type and size of b
	println()
	//Floating Point Type
	a3, b3 := 6.75, 7.5
	fmt.Printf("type of a %T b %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	//Complex Type
	c1 := complex(5, 7)
	c2 := 8 + 15i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	//String Data Type
	first := "Akshith"
	last := "Vanaparthy"
	name := first + " " + last
	fmt.Println("My name is", name)

	i := 55
	j := 55.6
	sum1 := i + int(j)
	fmt.Println(sum1)
}
