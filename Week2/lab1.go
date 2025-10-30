// Task 1: Write a Go program to demonstrate the use of assignment operators. The program should take two integer inputs and perform various assignment operations such as =, +=, -=, *=, /=, and %=. Display the result after each operation.
package main

import "fmt"

func main() {
	var num1, num2 int

	// Taking two integer inputs from the user
	fmt.Print("Enter first integer: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&num2)

	// Operation =
	c := num1
	// fmt.Printf("Answer Operation (=):%d\n", c)
	fmt.Printf("Operation (=):%d\n", c)

	// Operation +=
	num1 += num2
	fmt.Printf("Operation (+=):%d\n", num1)

	// Operation -=
	num1 -= num2
	fmt.Printf("Operation (-=):%d\n", num1)

	// Operation *=
	num1 *= num2
	fmt.Printf("Operation (*=):%d\n", num1)

	// Operation /=
	if num2 != 0 {
		num1 /= num2
		fmt.Printf("Operation (/=):%d\n", num1)
	} else {
		fmt.Println("Division by zero is not allowed for (/=) operation.")
	}

	// Operation %=
	if num2 != 0 {
		num1%= num2
		fmt.Printf("Operation (%%=):%d", num1)
	} else {
		fmt.Println("Division by zero is not allowed for (%%=) operation.")
	}
}	
