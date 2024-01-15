package main

import "fmt"

func main() {

	for i := 1; i <= 10; i += 1 {
		fmt.Printf(" %d", i)
	}

	for i := 1; i <= 10; i += 1 {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	for i := 1; i <= 10; i += 1 {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	n := 6
	for i := 0; i < n; i += 1 {
		for j := 0; j <= i; j += 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 3; i += 1 {
		for j := 1; j < 4; j += 1 {
			fmt.Printf("i = %d , j = %d\n", i, j)
		}

	}

	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 3
	}

}
