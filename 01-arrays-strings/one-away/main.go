package main

import (
	"fmt"
)

// there are three types of edits that can be performed on a string: insert, remove, and replace
// given two strings, write a function to determine if they are one edit away from each other

func main() {
	fmt.Println("Testing function oneAway")
	fmt.Print("'pale' & 'ple' -> ") // true
	fmt.Println(oneAway("pale", "ple"))
	fmt.Print("'pales' & 'pale' -> ") // true
	fmt.Println(oneAway("pales", "pale"))
	fmt.Print("'pale' & 'bale' -> ") // true
	fmt.Println(oneAway("pale", "bale"))
	fmt.Print("'pale' & 'bake' -> ") // false
	fmt.Println(oneAway("pale", "bake"))
	fmt.Print("'' & 'a' -> ") // true
	fmt.Println(oneAway("", "a"))
	fmt.Print("'hello' & 'world' -> ") // false
	fmt.Println(oneAway("hello", "world"))
	fmt.Print("'' & 'asdf' -> ") // false
	fmt.Println(oneAway("", "asdf"))
}

func oneAway(s1, s2 string) bool {

	if Abs(len(s1)-len(s2)) > 1 {
		return false
	}

	// count each char in s1
	count := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		count[s1[i]]++
	}

	// subtract each char in s2
	for i := 0; i < len(s2); i++ {
		count[s2[i]]--
	}

	// compare values to expected values to determine how many changes are needed to have equal strings
	hits := 0
	for _, value := range count {
		if Abs(value) == 1 {
			hits++
		} else if Abs(value) > 1 {
			return false
		}
	}

	if hits > 2 {
		return false
	}

	return true
}

// Abs is an int version of Abs
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
