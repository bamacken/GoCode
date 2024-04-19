package medium

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "babad", expected: "bab"},
		{input: "cbbd", expected: "bb"},
		{input: "a", expected: "a"},
		{input: "ac", expected: "a"},
	}

	for _, test := range tests {
		result := LongestPalindrome(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for input %s", test.expected, result, test.input)
		}
	}
}

func ExampleLongestPalindrome() {
	result := LongestPalindrome("babad")
	fmt.Println(result)
	// Output: bab
}

func BenchmarkLongestPalindrome(b *testing.B) {
	input := "babad"
	for i := 0; i < b.N; i++ {
		LongestPalindrome(input)
	}
}

func TestLongestPalindromeBruteForce(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "babad", expected: "bab"},
		{input: "cbbd", expected: "bb"},
		{input: "a", expected: "a"},
		{input: "ac", expected: "a"},
	}

	for _, test := range tests {
		result := LongestPalindromeBruteForce(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for input %s", test.expected, result, test.input)
		}
	}
}

func ExampleLongestPalindromeBruteForce() {
	result := LongestPalindromeBruteForce("babad")
	fmt.Println(result)
	// Output: bab
}

func BenchmarkLongestPalindromeBruteForce(b *testing.B) {
	input := "babad"
	for i := 0; i < b.N; i++ {
		LongestPalindromeBruteForce(input)
	}
}
