package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	price, no := 100, 7
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(6.5, 3.5)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	println()
	area2, _ := rectProps(10.4, 7.8)
	fmt.Printf("Area %f ", area2)
}
