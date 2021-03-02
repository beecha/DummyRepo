package main

import (
	"fmt"
)

// Convert a decimal representation of a number into
// Binary, Octal, and Hexal and display it
// NOTE: It's possible to do it by importing the "strconv"
// and using the ParseInt method, BUT please try to convert
// the number without using it
// BONUS: Try representing negative numbers
// BONUS: Write a function to check for overflow

func NumberRepresentation(x int) {
	fmt.Print("In binary: ")
	DecimalBinary(x)
	fmt.Print("\nIn hexadecimal: 0x")
	DecimalHex(x)

}

// From Decimal to Binary
func DecimalBinary(x int) {
	// Binary meaning base 2
	// Keep dividing the number by 2 until the quotient is zero
	// The remainder determines the bit
	// Example: 11
	// 11/2 = 5 remainder 1
	// 5/2 = 2 remainder 1
	// 2/2 = 1 remainder 0
	// 1/2 = 0 remainder 1
	// 11 in binary is 1011
	if x == 0 {
		fmt.Println(0)
	}

	result := []int{} // Empty int slice
	for x > 0 {
		result = append(result, (x % 2)) // append the remainder into the slice
		x = x / 2
	}

	// Print slice backwards
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}

func DecimalHex(x int) {
	// Hexadecimal is base 16
	// Keep dividing the number by 16 until the quotient is zero
	// Example: 250
	// 250 / 16 = 15 remainder 10
	// 15 / 16 = 0 remainder 15
	// result is "15""10" or 0xFA

	result := []int{}
	for x > 0 {
		result = append(result, (x % 16))
		x = x / 16
	}

	// We need to represent numbers 10 to 15 using characters A to F
	// A switch/case statement is perfect for it
	// Print slice backwards
	for i := len(result) - 1; i >= 0; i-- {
		switch result[i] {
		case 10:
			fmt.Print("A")
		case 11:
			fmt.Print("B")
		case 12:
			fmt.Print("C")
		case 13:
			fmt.Print("D")
		case 14:
			fmt.Print("E")
		case 15:
			fmt.Print("F")
		default:
			fmt.Print(result[i])
		}
	}

}

func main() {

	var userNumber int
	fmt.Print("Please enter a number:")
	fmt.Scanln(&userNumber)

	NumberRepresentation(userNumber)
}
