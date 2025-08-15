package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}
	s1 := greetings[:2]
	s2 := greetings[1:4]
	s3 := greetings[3:]
	fmt.Println(s1, s2, s3)
}

func exercise2() {
	message := "Hi 👩 and 👨"
	fmt.Println(string([]rune(message)[3]))
}

func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	fmt.Println(Employee{
		"Alice",
		"Smith",
		123,
	})

	fmt.Println(Employee{
		firstName: "Alice",
		lastName:  "Smith",
		id:        123,
	})

	var e Employee
	e.firstName = "Alice"
	e.lastName = "Smith"
	e.id = 123
	fmt.Println(e)
}
