package main

import (
	"fmt"
	"strings"
)

// given a string, write a function to check if it is a permutatino of a palindrome.
// the palindrome does not need to be limited to just dictionary words.
// you can ignrore casing and non-letter characters

func main() {
	fmt.Println("Testing function palindromePermutation")
	fmt.Print("String: 'Hello world' Output: ")
	fmt.Println(palindromePermutation("Hello world"))
	fmt.Print("String: 'Four words long string' Output: ")
	fmt.Println(palindromePermutation("Four words long string"))
	fmt.Print("String: 'Tact coa' Output: ")
	fmt.Println(palindromePermutation("Tact coa"))
	fmt.Print("String: 'racecar' Output: ")
	fmt.Println(palindromePermutation("racecar"))
	fmt.Print("String: '' Output: ")
	fmt.Println(palindromePermutation(""))
}

func palindromePermutation(s string) bool {

	if s == "" {
		return false
	}

	// strip all whitespace and convert to lowercase
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// iterate over the string and count the occurences of every letter
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	// check for valid palindrome conditions
	odds := false
	for _, value := range count {
		if value%2 == 1 && !odds {
			odds = true
		} else if value%2 == 1 && odds {
			return false
		}
	}

	return true
}
