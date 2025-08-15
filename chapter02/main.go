package main

import "fmt"

const value = 20

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	i := 20
	var f float64 = float64(i)

	fmt.Println(i, f)
}

func exercise2() {
	var i int = value
	var f float64 = float64(value)

	fmt.Println(i, f)
}

func exercise3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b, smallI, bigI)
	fmt.Println(b+1, smallI+1, bigI+1) // 0 -2147483648 0
}
