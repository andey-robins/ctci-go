package main

import (
	"fmt"
	"sort"
	"strings"
)

// given two strings, write a method to decide if one is a permutation of the other

func main() {
	fmt.Println("Testing function checkPermutation")
	fmt.Println("Checking strings: 'abcd' & 'dcba'")
	fmt.Println(checkPermutation("abcd", "dcba"))
	fmt.Println("Checking strings: 'aaaa' & 'bbbb'")
	fmt.Println(checkPermutation("aaaa", "bbbb"))
	fmt.Println("Checking strings: '' & 'dcba'")
	fmt.Println(checkPermutation("", "dcba"))
	fmt.Println("Checking strings: '' & ''")
	fmt.Println(checkPermutation("", ""))
}

// O (n log n)
func checkPermutation(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	splitOne := strings.Split(s1, "") // O(n log n)
	sort.Strings(splitOne)

	splitTwo := strings.Split(s2, "") // O(n log n)
	sort.Strings(splitTwo)

	// O(n)
	for i, _ := range splitOne {
		if splitOne[i] != splitTwo[i] {
			return false
		}
	}

	return true
}
