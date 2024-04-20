package hard

import (
	sort "GoCode/sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := 0.0
	// Check if the length of nums1 and nums2 are within the constraints
	if len(nums1) >= 0 && len(nums1) <= 1000 && len(nums2) >= 0 && len(nums2) <= 1000 {
		if len(nums1)+len(nums2) >= 1 && len(nums1)+len(nums2) <= 2000 {

			// Merge the two arrays
			merge := make([]int, len(nums1)+len(nums2))
			sort.MergeArray(merge, len(nums1), nums1, len(nums2), nums2)

			// Find the median of the sorted array
			mid := len(merge) / 2
			num1, num2 := 0, 0

			if len(merge)%2 != 0 {
				result = float64(merge[mid])
			} else {
				num1 = merge[mid-1]
				num2 = merge[mid]
				result = float64(num1+num2) / 2
			}
		}
	}

	return result
}

// sorts via merge sort
