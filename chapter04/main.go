package main

import (
	"fmt"
	"math/rand"
)

func main() {
	l := 100
	x := make([]int, 0, l)
	for range l {
		x = append(x, rand.Intn(100))
	}

	for _, xs := range x {
		switch {
		case xs%6 == 0:
			fmt.Println("Six!")
		case xs%2 == 0:
			fmt.Println("Fizz")
		case xs%3 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println("Never mind")
		}
	}

	var total int
	for i := range 10 {
		total := total + 1
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	// total をシャドーイングしているため、0 が出力される
	fmt.Printf("total=%v\n", total)
}
