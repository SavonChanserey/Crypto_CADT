package main

import ("fmt"; "encoding/hex"; "encoding/base64")

// text to binary
func stringToBin(input string) string {
	bin := ""
	for i := 0; i < len(input); i++ {
		bin += fmt.Sprintf("%08b ", input[i])
	}
	return bin
}

// text to hexadecimal


func main() {

	var input string

	fmt.Print("Enter a letter: ")
	fmt.Scan(&input)

	// Convert text to binary
	bin := stringToBin(input)
	fmt.Printf("The binary representation of '%s' is: %s\n", input, bin)

	// Convert text to hexadecimal
	hex := hex.EncodeToString([]byte(input))
	fmt.Printf("The hexadecimal representation of '%s' is: %s\n", input, hex)

	// Convert text to Base64
	base64 := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Printf("The Base64 representation of '%s' is: %s\n", input, base64)

}