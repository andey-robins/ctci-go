package main

import (
	"fmt"
	"strconv"
)

// implement a method to perform basic string compression using the count of repeated characters
// aabcccccaaa -> a2b1c5a3
// if the "compressed" string would not be shorter, return the original string
// you can assume only uppercase and lowercase letters

func main() {
	fmt.Println("Testing function stringCompression")
	fmt.Print("String: 'llll' Expected Output: 'l4' Output: ")
	fmt.Println(stringCompression("llll"))
	fmt.Print("String: 'Hello' Expected Output: 'Hello' Output: ")
	fmt.Println(stringCompression("Hello"))
	fmt.Print("String: 'aabcccccaaa' Expected Output: 'a2b1c5a3' Output: ")
	fmt.Println(stringCompression("aabcccccaaa"))
}

func stringCompression(s string) string {

	out := ""
	last := s[0]
	count := 0

	for i, _ := range s {
		if s[i] == last {
			count++
		} else {
			out += string(last)
			out += strconv.Itoa(count)
			count = 1
			last = s[i]
		}

		// return early if the compressed string is going to be longer than the original
		if len(out) >= len(s) {
			return s
		}
	}

	// output the last tracked character
	out += string(last)
	out += strconv.Itoa(count)
	return out
}
