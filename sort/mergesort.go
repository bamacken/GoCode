package sort

// The merge portion of a mergesort which takes 2 sorted arrays and returns one sorted array
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

// MergeSort sorts an array using the merge sort algorithm
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	for i := 0; i < mid; i++ {
		left[i] = arr[i]
	}

	for i := mid; i < len(arr); i++ {
		right[i-mid] = arr[i]
	}

	MergeSort(left)
	MergeSort(right)
	merge(left, right, arr)
}

func merge(left, right, arr []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
