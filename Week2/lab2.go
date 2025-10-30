package main

import "fmt"

func main() {
	var num1, num2 int

	// Taking two integer inputs from the user
	fmt.Print("Enter first integer (a): ")
	fmt.Scan(&num1)
	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&num2)

	// both positive (&&)
	positive := (num1 > 0) && (num2 > 0)
	greater := (num1 > num2) || (num2 > num1)
	notequal := !(num1 == num2)
	
	fmt.Printf("Both numbers are positive: %t\n", positive)
	fmt.Printf("At least one number is greater: %t\n", greater)
	fmt.Printf("Numbers are not equal: %t\n", notequal)
}