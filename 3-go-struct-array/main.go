package main

import "fmt"

// create Student struct
type Student struct {
	id      int
	age     int
	address string
}

// print console students info
func displayInfo(name string, student Student) {
	fmt.Printf("%v : Id = %v, Age = %v, Address = %v\n", name, student.id, student.age, student.address)
}

// change struct value
func (s *Student) increaseAge(val int) {
	s.age = s.age + val
}

func main() {
	mukul := Student{101, 24, "Mymensingh"}
	alam := Student{102, 26, "Mymensingh"}

	displayInfo("Mukul", mukul)
	displayInfo("Alam", alam)

	// called increaseAge function to change struct value
	mukul.increaseAge(6)
	displayInfo("Mukul", mukul)
}
