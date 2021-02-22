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
	for i := 0; i < len(s); i++ {
		result = string(s[i]) + result
	}

	return result
}

func ReverseString2(s string) string {
	var result []byte // Each character in a string are bytes

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return string(result)
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
	testCases := []string{
		"racecar",
		"To be, or not",
		"ABC123",
		"qwerrewq",
		".     .",
		"is it not so is it",
	}

	for _, v := range testCases {
		fmt.Printf("%v\n%v\nPalindrome: %v\n", v, ReverseString(v), IsPalindrome(v))
	}
}
