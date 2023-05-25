package main

import "testing"

func TestIsUnique(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"abcdefg", true},
		{"aaaa", false},
		{"abab", false},
		{"", true},
	}

	for _, tt := range tests {
		actual := isUnique(tt.input)
		if actual != tt.expected {
			t.Errorf("isUnique(%s): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
}
