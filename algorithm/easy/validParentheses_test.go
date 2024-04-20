package easy

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "([)]", expected: false},
		{input: "{[]}", expected: true},
	}

	// Run tests
	for _, tc := range testCases {
		actual := IsValid(tc.input)
		if actual != tc.expected {
			t.Errorf("Input: %s, Expected: %v, Got: %v", tc.input, tc.expected, actual)
		}
	}
}
func BenchmarkIsValid(b *testing.B) {
	input := "({[]})" // Example input
	for i := 0; i < b.N; i++ {
		IsValid(input)
	}
}

func ExampleIsValid() {
	fmt.Println(IsValid("{[]}")) // Output: true
}
