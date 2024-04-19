package medium

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	// Test case 1: Positive number
	input1 := 123
	expected1 := 321
	output1 := Reverse(input1)
	if output1 != expected1 {
		t.Errorf("Reverse(%d) = %d; want %d", input1, output1, expected1)
	}

	// Test case 2: Negative number
	input2 := -456
	expected2 := -654
	output2 := Reverse(input2)
	if output2 != expected2 {
		t.Errorf("Reverse(%d) = %d; want %d", input2, output2, expected2)
	}

	// Test case 3: Single digit number
	input3 := 7
	expected3 := 7
	output3 := Reverse(input3)
	if output3 != expected3 {
		t.Errorf("Reverse(%d) = %d; want %d", input3, output3, expected3)
	}
}

// ExampleReverse demonstrates the usage of the Reverse function.
func ExampleReverse() {
	input := 123
	output := Reverse(input)
	fmt.Println(output)
	// Output: 321
}

// BenchmarkReverse measures the performance of the Reverse function.
func BenchmarkReverse(b *testing.B) {
	input := 123
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}
