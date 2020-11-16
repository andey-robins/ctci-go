package main

import (
	"fmt"
	"strings"
)

// Write a method to replace all spaces in a string with %20.
// you may assume that the string has enough space at the end to store all the additional characters
// and that you are given the "true" length of the string.

func main() {
	fmt.Println("Testing function urlify")
	fmt.Print("String: 'Hello world' Output: ")
	fmt.Println(urlify("Hello world  ", 10))
	fmt.Print("String: 'Four words long string' Output: ")
	fmt.Println(urlify("Four words long string      ", 19))
}

// this isn't an in place solution, but it is going to be the most readable, simmplest version possible
func urlify(s string, i int) string {
	return strings.Replace(strings.TrimSpace(s), " ", "%20", -1)
}
