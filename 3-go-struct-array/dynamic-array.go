package main

import "fmt"

func main() {
	// dynamic array (slice) declare
	var students []string
	var length int
	var name string

	// get user input to array size
	fmt.Printf("How many students data are inserted: ")
	fmt.Scan(&length)

	// get array elements value form user
	for i := 0; i < length; i++ {
		fmt.Printf("Enter %v name: ", i+1)
		fmt.Scan(&name)
		students = append(students, name)
	}

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	fmt.Printf("Total students = %v", len(students))
}
