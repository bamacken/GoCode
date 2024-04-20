package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeArray(t *testing.T) {
	// Test cases
	tests := []struct {
		name   string
		merge  []int
		len1   int
		nums1  []int
		len2   int
		nums2  []int
		result []int
	}{
		{
			name:   "Example 1",
			merge:  make([]int, 7),
			len1:   4,
			nums1:  []int{1, 3, 5, 7},
			len2:   3,
			nums2:  []int{2, 4, 6},
			result: []int{1, 2, 3, 4, 5, 6, 7},
		},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeArray(tt.merge, tt.len1, tt.nums1, tt.len2, tt.nums2)
			if !reflect.DeepEqual(tt.merge, tt.result) {
				t.Errorf("mergeArray() = %v, want %v", tt.merge, tt.result)
			}
		})
	}
}

func ExampleMergeArray() {
	merge := make([]int, 7)
	len1 := 4
	nums1 := []int{1, 3, 5, 7}
	len2 := 3
	nums2 := []int{2, 4, 6}

	MergeArray(merge, len1, nums1, len2, nums2)
	fmt.Println(merge)
	// Output: [1 2 3 4 5 6 7]
}

func BenchmarkMergeArray(b *testing.B) {
	merge := make([]int, 7)
	len1 := 4
	nums1 := []int{1, 3, 5, 7}
	len2 := 3
	nums2 := []int{2, 4, 6}

	for i := 0; i < b.N; i++ {
		MergeArray(merge, len1, nums1, len2, nums2)
	}
}
