package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindBuildings(t *testing.T) {
	// Test case 1
	heights := []int{1, 3, 2, 4}
	expected := []int{3}
	result := FindBuildings(heights)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 2
	heights = []int{4, 2, 3, 1}
	expected = []int{0, 2, 3}
	result = FindBuildings(heights)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// ExampleFindBuildings demonstrates the usage of the FindBuildings function.
func ExampleFindBuildings() {
	heights := []int{1, 3, 2, 4}
	result := FindBuildings(heights)
	fmt.Println(result)
	// Output: [3]
}

// BenchmarkFindBuildings measures the performance of the FindBuildings function.
func BenchmarkFindBuildings(b *testing.B) {
	heights := []int{1, 3, 2, 4}
	for i := 0; i < b.N; i++ {
		FindBuildings(heights)
	}
}
