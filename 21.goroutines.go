package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(750 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func main() {
	go numbers()
	time.Sleep(5100 * time.Millisecond)
	fmt.Println("main terminated")
}
