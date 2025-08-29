package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

func exercise1() {
	person := MakePerson("Radish", "Miyazaki", 31)
	personPointer := MakePersonPointer("Radish", "Miyazaki", 31)

	fmt.Println(person)
	fmt.Println(personPointer)
}

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println(ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}

func exercise2() {
	// UpdateSlice はスライスを受け取り、その要素を直接変更しているため、呼び出し元のスライスも変更する
	ss := []string{"a", "b", "c"}
	fmt.Printf("UpdateSlice を呼び出す前: %v\n", ss)
	UpdateSlice(ss, "d")
	fmt.Printf("UpdateSlice を呼び出した後: %v\n", ss)

	// GrowSlice はスライスを受け取り、新しいスライスを追加しているが、
	// len が変更されるのは GrowSlice の中のみであるため、呼び出し元のスライスは変更されない
	ss = []string{"a", "b", "c"}
	fmt.Printf("GrowSlice を呼び出す前: %v\n", ss)
	GrowSlice(ss, "d")
	fmt.Printf("GrowSlice を呼び出した後: %v\n", ss)
}

func exercise3() {
	var people []Person
	for range 10_000_000 {
		people = append(people, Person{
			FirstName: "Radish",
			LastName:  "Miyazaki",
			Age:       31,
		})
	}
}

func exercise4() {
	people := make([]Person, 10_000_000)
	for i := range people {
		people[i] = Person{
			FirstName: "Radish",
			LastName:  "Miyazaki",
			Age:       31,
		}
	}
}
