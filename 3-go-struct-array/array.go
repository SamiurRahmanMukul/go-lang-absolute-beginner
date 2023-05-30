package main

import "fmt"

func main() {
	// array declare and initialize
	var destination = [5]string{
		"Bangladesh", "Pakistan", "India", "Nepal", "Australia",
	}

	for i := 0; i < len(destination); i++ {
		fmt.Println(destination[i])
	}

	// array declare
	var name [5]string

	// get array elements value form user
	for i := 0; i < len(name); i++ {
		fmt.Printf("Enter %v name: ", i+1)
		fmt.Scan(&name[i])
	}

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}
