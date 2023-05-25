// implement an algorithm to determine if a string has all unique characters
// what if you cannot use additional data structures?

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Testing function isUnique")
	fmt.Println("String: 'abcdefg'")
	fmt.Println(isUnique("abcdefg"))
	fmt.Println("String: 'aaaa'")
	fmt.Println(isUnique("aaaa"))
	fmt.Println("String: 'abab'")
	fmt.Println(isUnique("abab"))
	fmt.Println("String: ''")
	fmt.Println(isUnique(""))
}

func isUnique(s string) bool {

	splitS := strings.Split(s, "")
	sort.Strings(splitS)

	for i, c := range splitS {
		if i < len(s)-1 && c == splitS[i+1] {
			return false
		}
	}

	return true
}
