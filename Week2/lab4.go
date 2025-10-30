package main

import ("fmt"; "strings")

func Addtion(num1, num2 int) int {
	result := num1 + num2
	fmt.Printf("Result: %d + %d = %d", num1, num2, result)
	return result
}

func Subtraction(num1, num2 int) int {
	result := num1 - num2
	fmt.Printf("Result: %d + %d = %d", num1, num2, result)
	return result
}

func Multiplication(num1, num2 int) int {
	result := num1 * num2
	fmt.Printf("Result: %d + %d = %d", num1, num2, result)
	return result
}

func Division(num1, num2 int) int {

	if num2 == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	result := num1 / num2
	fmt.Printf("Result: %d + %d = %d", num1, num2, result)
	return result
}

func Modulus(num1, num2 int) int {

	if num2 == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	result := num1 % num2
	fmt.Printf("Result: %d + %d = %d", num1, num2, result)
	return result
}


func main() {
	var num1, num2, operator int

	fmt.Print(strings.Repeat("=",16))
	fmt.Print("Mini Calculator")
	fmt.Print(strings.Repeat("=",16))

	fmt.Println()

	fmt.Print("1. Addtion\n2. Substract\n3. Multiplication\n4. Division\n5. Modulus\n6. Exit\n")

	fmt.Print("Please choose operator:")
	fmt.Scan(&operator)

	if (operator == 6) {
		fmt.Println("Exiting the program.")
		return
	}

	fmt.Print("Enter the first integer:")
	fmt.Scan(&num1)

	fmt.Print("Enter the second integer:")
	fmt.Scan(&num2)

	switch operator { 
		case 1:
			Addtion(num1, num2)
		case 2:
			Subtraction(num1, num2)
		case 3:
			Multiplication(num1, num2)
		case 4:
			Division(num1, num2)
		case 5:
			Modulus(num1, num2)
		default:
			fmt.Println("Invalid operator selected.")
			return
	}

}