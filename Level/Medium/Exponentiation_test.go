package medium

import (
	"fmt"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	// Test case 1: Positive exponent
	x1 := 2.0
	n1 := 3
	expected1 := 8.0
	result1 := MyPow(x1, n1)
	if math.Abs(result1-expected1) > 1e-9 {
		t.Errorf("Test case 1 failed: expected %f, got %f", expected1, result1)
	}

	// Test case 2: Negative exponent
	x2 := 3.0
	n2 := -2
	expected2 := 1.0 / 9.0
	result2 := MyPow(x2, n2)
	if math.Abs(result2-expected2) > 1e-9 {
		t.Errorf("Test case 2 failed: expected %f, got %f", expected2, result2)
	}

	// Test case 3: Zero exponent
	x3 := 5.0
	n3 := 0
	expected3 := 1.0
	result3 := MyPow(x3, n3)
	if math.Abs(result3-expected3) > 1e-9 {
		t.Errorf("Test case 3 failed: expected %f, got %f", expected3, result3)
	}
}

// ExampleMyPow demonstrates the usage of the MyPow function.
func ExampleMyPow() {
	x := 2.0
	n := 4
	result := MyPow(x, n)
	fmt.Printf("%.1f\n", result)
	// Output: 16.0
}

// BenchmarkMyPow measures the performance of the MyPow function.
func BenchmarkMyPow(b *testing.B) {
	x := 2.0
	n := 10
	for i := 0; i < b.N; i++ {
		MyPow(x, n)
	}
}
