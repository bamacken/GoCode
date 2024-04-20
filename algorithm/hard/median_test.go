package hard

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	// Test cases
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		// Add more test cases here
	}

	for _, test := range tests {
		result := FindMedianSortedArrays(test.nums1, test.nums2)
		if result != test.expected {
			t.Errorf("Expected %f, but got %f", test.expected, result)
		}
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 4, 6, 8, 10}

	for i := 0; i < b.N; i++ {
		FindMedianSortedArrays(nums1, nums2)
	}
}
