package main

import (
	"fmt"
)

func xorEncrypt(text string, key byte) string { 

	encrypted := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encrypted[i] = text[i] ^ key
	}
	return string(encrypted)
}

func main() {

	var input string
	var key byte

	fmt.Printf("Enter the strings:")
	fmt.Scan(&input)

	fmt.Printf(("Enter the key:"))
	fmt.Scan(&key)
	
	// Encrypt
	EncryptedText := xorEncrypt(input, key)
	fmt.Printf("Encrypted Text: %s\n", EncryptedText)

	// Decrypt
	DecryptedText := xorEncrypt(EncryptedText, key)
	fmt.Printf("Decrypted Text: %s\n", DecryptedText)

}
