package main

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	var tests = []struct {
		inputOne string
		inputTwo string
		expected bool
	}{
		{"abcd", "dcba", true},
		{"aaaa", "bbbb", false},
		{"", "dcba", false},
		{"", "", true},
	}

	for _, tt := range tests {
		actual := checkPermutation(tt.inputOne, tt.inputTwo)
		if actual != tt.expected {
			t.Errorf("checkPermutation(%s, %s): expected %t, actual %t", tt.inputOne, tt.inputTwo, tt.expected, actual)
		}
	}
}
