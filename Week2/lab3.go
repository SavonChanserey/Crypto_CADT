package main

import (
	"fmt"
)

func MyXOR(num1, num2 int){
	xor := num1 ^ num2
	fmt.Printf("XOR(%d,%d) = %04b = %d\n" ,num1 ,num2 ,xor ,xor)
}

func MyNOT(num1 int){
	not := ^num1
	fmt.Printf("NOT(%d) = %04b = %d\n",num1 ,not, not)
}

func MyOR(num1, num2 int){
	or := num1 | num2
	fmt.Printf("OR(%d,%d) =  %04b = %d\n",num1 ,num2 ,or , or)
}

func MyAND(num1, num2 int){
	and := num1 & num2
	fmt.Printf("AND(%d,%d) =  %04b = %d\n",num1 ,num2 ,and, and)
}

func LeftShift(num1 int, n uint){
	left := num1 << n
	fmt.Printf("LeftShift(%d,%d) =  %04b = %d\n", num1, n, left, left)
}

func RightShift(num1 int, n uint){
	right := num1 >> n
	fmt.Printf("RightShift(%d,%d) = %04b = %d\n", num1, n, right, right)
}

// %b convert to binary 

func main(){
	
	var num1 = 8
	var num2 = 9

	MyXOR(num1, num2)
	MyNOT(num1)
	MyOR(num1, num2)
	MyAND(num1, num2)
	LeftShift(num1, 1)
	RightShift(num1, 1)
	
}
