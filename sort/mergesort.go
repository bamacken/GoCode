package sort

func MergeArray(merge []int, len1 int, nums1 []int, len2 int, nums2 []int) {
	i, j, k := 0, 0, 0
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			merge[k] = nums1[i]
			i++
		} else {
			merge[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len1 {
		merge[k] = nums1[i]
		i++
		k++
	}

	for j < len2 {
		merge[k] = nums2[j]
		j++
		k++
	}
}
