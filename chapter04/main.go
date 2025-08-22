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
		if xs%2 == 0 || xs%3 == 0 {
			fmt.Println("Six!")
		} else if xs%2 == 0 {
			fmt.Println("Fizz")
		} else if xs%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println("Never mind")
		}
	}

	var total int
	for i := range 10 {
		total := total + 1
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("total=%v\n", total)
}
