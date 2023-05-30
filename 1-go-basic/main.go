package main

import "fmt"

func main() {
	// TODO: run go first program
	// fmt.Printf("Hello World! \n")

	/*
		TODO: Token of any programming language
		1. keywords
		2. data types
		3. variables
		4. escapes sequences
		5. operators
	*/

	/*
		TODO: data types of go programming language
		1. boolean
		2. string
		3. numeric - integer, floating
		4. derived - pointer, array, structure, slice, map
	*/

	/*
		TODO: variable declare rules
		1. letters, digits, _
		2. keywords are not allowed as variable
		3. variable can not have space
		4. variable name can not be start with digits
	*/

	// TODO: variable declaration
	/* var name, countryName string
	var age int
	var gpa float32 */

	// TODO: variable initialization
	/* name = "Samiur Rahman Mukul"
	countryName = "Bangladesh"
	age = 24
	gpa = 4.00 */

	// TODO: variable usages
	/* fmt.Println("My name is ", name)
	fmt.Println("I am", age, "years old")
	fmt.Println("My county name is", countryName)
	fmt.Println("I got", gpa, "GPA in HSC exam") */

	// TODO: shortcut variable declare & initialize
	/* callName := "Mukul"
	fmt.Println(callName) */

	// TODO: constant variable & formatting output
	/* const COUNTRY = "Bangladesh"
	fmt.Printf("%v is %v years old and he live in %v", name, age, COUNTRY) */

	// TODO: getting user input
	// input -> Scan, Scanln, Scanf
	// output -> Print, Println, Printf

	/* var firstName, lastName string

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName) // ? getting input in same line
	fmt.Scan(&firstName) // ? getting input in new lint

	fmt.Print("\nPlease enter your first and last name: ")
	fmt.Scanf("%v %v", &firstName, &lastName)
	fmt.Printf("Your full name is %v %v", firstName, lastName) */

	// TODO: number formatting
	/* var decimalNumber int

	fmt.Printf("Please enter decimal number: ")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("\n Binary number = %b", decimalNumber)
	fmt.Printf("\n Octal number = %o", decimalNumber)
	fmt.Printf("\n Hex number = %x", decimalNumber) */

	// TODO: string formatting
	/* name := "Md. Samiur Rahman (Mukul)"
	number := 3.1416

	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)
	fmt.Printf("%f\n", number)
	fmt.Printf("%.2f\n", number) */

	/*
		TODO: types of operators
		 1. Arithmetic Operators -> +, , *, /, %
		 2. Assignment Operators -> =, +=, -=, *=, /=, %=
		 3. Unary Operators -> ++, --
		 4. Relational Operators -> ==, !=, <, >, <=, >=
		 5. Logical Operators -> &&, ||, !
		 6. Bitwise Operators -> &, |, ^, >>, <<
		 7. Special Operators
	*/

	// TODO: arithmetic operators -> +, , *, /, %
	/* var num1, num2 int
	fmt.Printf("Enter 2 number: ")
	fmt.Scanf("%v %v", &num1, &num2)

	fmt.Printf("\n%v + %v = %v", num1, num1, num1+num2)
	fmt.Printf("\n%v - %v = %v", num1, num1, num1-num2)
	fmt.Printf("\n%v * %v = %v", num1, num1, num1*num2)
	// type conversation -> int to float
	fmt.Printf("\n%v / %v = %.3f", num1, num1, float32(num1)/float32(num2))
	fmt.Printf("\n%v %% %v = %v", num1, num1, num1%num2) */

	// TODO: problem solve 1 -> Area calculate of circle & rectangle
	// 1.1: area of triangle = 0.5 * base * height
	/* var base, height float32
	fmt.Printf("Enter Base and Height: ")
	fmt.Scanf("%v %v %v", &base, &height)

	area := 0.5 * base * height
	fmt.Printf("\nArea of Triangle = %v", area) */

	// 1.2: area of circle = 3.1416 * radius * radius
	/* var radius float32
	fmt.Printf("\nEnter radius: ")
	fmt.Scanf("%v", &radius)

	areaOfCircle := math.Pi * radius * radius
	fmt.Printf("\nArea of Circle = %v", areaOfCircle) */

	// TODO: assignment operator -> =, +=, -=, *=, /=, %=
	/* x := 4
	x = x + 5
	x += 5 // x = x + 5
	x -= 5 // x = x - 5
	x *= 5 // x = x * 5
	x /= 5 // x = x / 5
	x %= 5 // x = x % 5
	fmt.Printf("X = %v", x) */

	// TODO: unary operator -> ++, --
	/* x := 5
	x++ // x = x + 5
	x-- // x = x - 5 */

	// TODO: relational operator -> ==, !=, <, >, <=, >=
	/* y := 5
	z := 6

	fmt.Printf("Y == Z = %v\n", y == z)
	fmt.Printf("Y != Z = %v\n", y != z)
	fmt.Printf("Y < Z = %v\n", y < z)
	fmt.Printf("Y > Z = %v\n", y > z)
	fmt.Printf("Y <= Z = %v\n", y <= z)
	fmt.Printf("Y >= Z = %v\n", y >= z) */

	// TODO: logical operator -> &&, ||, !
	/* fmt.Printf("3>4 && 5>6 = %v\n", 3>4 && 5>6)
	fmt.Printf("3>4 || 5>6 = %v\n", 3>4 && 5>6)
	fmt.Printf("!(5>6) = %v\n", !(5>6)) */

	/*
		TODO: if, else if, else statement
		1: control statement - conditional control statement, loop control statement
		1.1: conditional control statement -> if, else if, else, switch
		1.2: loop control statement -> for loop
	*/

	// TODO: if, else if, else
	/* const number = 10

	if number > 0 {
		fmt.Printf("Number is Positive!")
	} else if number < 0 {
		fmt.Printf("Number is Negative!")
	} else if number == 0 {
		fmt.Printf("Number is Zero!")
	} else {
		fmt.Printf("Unknown error")
	} */

	// TODO: problem solve 2 -> Event or Odd number detect
	/* var num int

	fmt.Printf("\nPlease enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("%v number is Even!", num)
	} else {
		fmt.Printf("%v number is Odd!", num)
	} */

	// TODO: problem solve 3 -> Largest number find out between 2 numbers
	/* var num1, num2 int

	fmt.Printf("\nPlease enter two number: ")
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Printf("%v is larger then %v!", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%v is larger then %v!", num2, num1)
	} else {
		fmt.Printf("%v is equal to %v!", num1, num2)
	} */

	// TODO: problem solve 4 -> Largest number find out between 3 numbers
	/* var number1, number2, number3 int

	fmt.Printf("\nPlease enter three number: ")
	fmt.Scan(&number1, &number2, &number3) */

	/* if number1 > number2 {
		if number1 > number3 {
			fmt.Printf("Larger number is %v", number1)
		} else {
			fmt.Printf("Larger number is %v", number3)
		}
	} else if number2 > number1 {
		if number2 > number3 {
			fmt.Printf("Larger number is %v", number2)
		} else {
			fmt.Printf("Larger number is %v", number3)
		}
	} else {
		fmt.Printf("Numbers are equal")
	} */

	/* if number1 > number2 && number1 > number3 {
		fmt.Printf("is the number\n", number1)
	} else if number2 > number1 && number2 > number3 {
		fmt.Printf("%v is the largest number\n", number2)
	} else if number3 > number1 && number3 > number2 {
		fmt.Printf("%v is the largest number\n", number3)
	} else {
		fmt.Printf("Numbers are equal\n")
	} */

	// TODO: problem solve 5 -> Digit spelling program using switch case
	/* var number int
	fmt.Printf("\nPlease enter a number: ")
	fmt.Scan(&number)

	switch number {
	case 0:
		fmt.Printf("Zero")
	case 1:
		fmt.Printf("One")
	case 2:
		fmt.Printf("Two")
	case 3:
		fmt.Printf("Three")
	case 4:
		fmt.Printf("Four")
	case 5:
		fmt.Printf("Five")
	case 6:
		fmt.Printf("Six")
	case 7:
		fmt.Printf("Seven")
	case 8:
		fmt.Printf("Eight")
	case 9:
		fmt.Printf("Nine")
	default:
		fmt.Printf("Not a valid digit")
	} */

	// TODO: problem solve 6 -> Multiple case in switch statement
	/* var classNumber int
	fmt.Printf("\nEnter your class: ")
	fmt.Scan(&classNumber)

	switch classNumber {
	case 1, 2, 3, 4, 5:
		fmt.Printf(" Primary\n ")
	case 6, 7, 8, 9, 10:
		fmt.Printf("Secondary\n")
	case 11, 12:
		fmt.Printf("Higher-Secondary\n")
	default:
		fmt.Printf("Not a valid digit\n")
	} */

	// TODO: problem solve 7 -> print even/odd numbers from 1 to 100 using for loop
	/* for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even number\n", i)
		} else {
			fmt.Printf("%v is odd number\n", i)
		}
	} */

	// TODO: problem solve 8 -> series --> 1 + 2 + 3 + ... + 10
	/* sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Printf("1 + 2 + 3 + ... + 10 series sum = %v\n", sum) */

	// TODO: break, continue statement
	/* for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		} else if i == 10 {
			break
		} else {
			fmt.Printf("%v\t", i)
		}
	} */

	// TODO: pointer - call by reference
	var p *int
	x := 10
	p = &x

	// p -> x memory address --> 10
	*p = *p - 1 // call by reference value

	fmt.Printf("&x 	= %v\n", &x) // -> x memory address
	fmt.Printf("p 	= %v\n", p)   // -> x memory address
	fmt.Printf("*p 	= %v\n", *p) // -> x value
}
