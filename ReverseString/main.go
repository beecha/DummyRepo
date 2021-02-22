package main

import (
	"fmt"
)

/*
	Given a string, write a function that takes a string
	and returns the string in reverse order

	Example:
	Input: Hello World!
	Output: !dlroW olleH
*/

func ReverseString(s string) string {
	result := "" // Start off with an emtpy string to append to
	n := len(s)

	// Grab each rune from the string and append
	// it to reuslt in reverse order
	for i := n; i >= 0; i-- {
		result.append(result, s[i])
	}

	return result
}

func main() {
	// Some Example Strings
	firstString := "I'll be back"
	/*
		secondString := "To be, or not to be"
		thirdString := "ABC123 321CBA"
		fourthString := "!@#$%^"
		fifthString := ".          !"
	*/
	fmt.Println(firstString[0])
	fmt.Printf("%c\n", 73)
	fmt.Println(ReverseString(firstString))
}
