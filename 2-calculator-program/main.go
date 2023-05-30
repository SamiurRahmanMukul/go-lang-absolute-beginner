package main

import "fmt"

func add(num1, num2 float32) float32 {
	return num1 + num2
}
func sum(num1, num2 float32) float32 {
	return num1 - num2
}
func mul(num1, num2 float32) float32 {
	return num1 * num2
}
func div(num1, num2 float32) float32 {
	return num1 / num2
}

func main() {
	var num1, num2, result float32
	var option string
	var exit int
	i := true

	for i {
		fmt.Printf("Enter num1 = ")
		fmt.Scan(&num1)
		fmt.Printf("Enter num2 = ")
		fmt.Scan(&num2)
		fmt.Printf("Chose an options (+, -, *, /) : ")
		fmt.Scan(&option)

		switch option {
		case "+":
			result = add(num1, num1)
		case "-":
			result = sum(num1, num1)
		case "*":
			result = mul(num1, num1)
		case "/":
			result = div(num1, num1)
		default:
			fmt.Println("Invalid Options")
			continue
		}

		fmt.Printf("Result = %v\n", result)

		// continue program or exit
		fmt.Printf("Are continue program to Enter: 1 or Exit to press any key and enter: ")
		fmt.Scan(&exit)

		if exit != 1 {
			break
		}
	}
}
