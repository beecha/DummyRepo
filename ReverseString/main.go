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

	Write a function that returns true if the string
	is a palindrome
*/

func ReverseString(s string) string {
	result := "" // Start off with an emtpy string to append to

	// Grab each character, typecast it as a string
	// then concatenate it to the resulting string
	for i := len(s) - 1; i >= 0; i-- {
		result = string(s[i]) + result
	}

	return result
}

func IsPalindrome(s string) bool {
	result := ReverseString(s)
	if result == s {
		return true
	}

	return false
}

func main() {
	// Some Example Strings
	firstString := "I'll be back"
	secondString := "To be, or not to be"
	thirdString := "ABC123 321CBA"
	fourthString := "!@#$%^"
	fifthString := ".          !"

	fmt.Println(ReverseString(firstString))
	fmt.Println(ReverseString(secondString))
	fmt.Println(ReverseString(thirdString))
	fmt.Println(ReverseString(fourthString))
	fmt.Println(ReverseString(fifthString))

	fmt.Println(IsPalindrome(thirdString))
}
