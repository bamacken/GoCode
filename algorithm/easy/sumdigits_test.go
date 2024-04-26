package easy

import "testing"

func TestSolution(t *testing.T) {
	// Test case 1: n = 0
	result := SumIntegerDigits(0)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: n = 12
	result = SumIntegerDigits(12)
	expected = 3
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 3: n = 99
	result = SumIntegerDigits(99)
	expected = 18
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 4: n = 37
	result = SumIntegerDigits(37)
	expected = 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
